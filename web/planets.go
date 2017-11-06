package web

import (
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/antihax/goesi/esi"
	"github.com/curzonj/eve-dwh-golang/model"
)

type PlanetData struct {
	Account          string
	CharacterID      int64
	Character        string
	PlanetName       string
	ConstelationName string
	PlanetType       string
	BIFCount         int
	Extracted        int
	NextAttention    time.Time
}

var (
	aifTypes       = []int32{2470, 2472, 2474, 2480, 2484, 2485, 2491, 2494}
	launchpadTypes = []int32{2256, 2542, 2543, 2544, 2552, 2555, 2556, 2557}
	storageTypes   = []int32{2257, 2535, 2536, 2541, 2558, 2560, 2561, 2562}
	bifTypes       = []int32{2469, 2471, 2473, 2481, 2483, 2490, 2492, 2493}
	extractorTypes = []int32{2848, 3060, 3061, 3062, 3063, 3064, 3067, 3068}
)

func intArrayContains(list []int32, a int32) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (h *handler) planets(w http.ResponseWriter, r *http.Request) error {
	session := session(r)
	userID := session.Values["user_id"].(string)
	logger := logger(r)

	var characters []model.UserCharacter
	err := h.clients.DB.Select(&characters, "select * from user_characters where user_id = $1", userID)
	if err != nil {
		return err
	}

	bc := make(chan PlanetData)
	var wg sync.WaitGroup

	for _, c := range characters {
		wg.Add(1)
		go func(c model.UserCharacter) {
			defer wg.Done()

			ctx, err := c.TokenSourceContext(r.Context(), h.clients)
			if err != nil {
				logger.Error(err)
				return
			}

			data, _, err := h.clients.EVEBreakerClient.ESI.PlanetaryInteractionApi.GetCharactersCharacterIdPlanets(ctx, int32(c.ID), nil)
			if err != nil {
				logger.Error(err)
				return
			}

			for _, j := range data {
				wg.Add(1)
				go func(j esi.GetCharactersCharacterIdPlanets200Ok) {
					defer wg.Done()

					data, _, err := h.clients.EVEBreakerClient.ESI.PlanetaryInteractionApi.GetCharactersCharacterIdPlanetsPlanetId(ctx, int32(c.ID), j.PlanetId, nil)
					if err != nil {
						logger.Error(err)
						return
					}

					count := 0
					extracted := 0
					NextAttention := time.Now().Add(time.Hour * time.Duration(10000))
					pinMap := make(map[int64]esi.GetCharactersCharacterIdPlanetsPlanetIdPin, len(data.Pins))
					routeMap := make(map[int64]map[int64]esi.GetCharactersCharacterIdPlanetsPlanetIdRoute)

					for _, r := range data.Routes {
						if routeMap[r.SourcePinId] == nil {
							routeMap[r.SourcePinId] = make(map[int64]esi.GetCharactersCharacterIdPlanetsPlanetIdRoute)
						}
						routeMap[r.SourcePinId][r.DestinationPinId] = r
					}

					for _, pin := range data.Pins {
						pinMap[pin.PinId] = pin
					}

					for _, pin := range data.Pins {
						if intArrayContains(bifTypes, pin.TypeId) {
							count = count + 1
						}

						if intArrayContains(extractorTypes, pin.TypeId) {
							extracted = extracted + int(pin.ExtractorDetails.QtyPerCycle)
							if pin.ExpiryTime.Before(NextAttention) {
								NextAttention = pin.ExpiryTime
							}
						}

						// Deadline calculation for AIFs
						// Logic shortcuts:
						// * AIFs
						// * Only one schematic per launchpad
						// * All the same amount of stuff per cycle for each input
						if intArrayContains(launchpadTypes, pin.TypeId) {
							srcMap := routeMap[pin.PinId]
							ratePerHour := int64(0)
							fewestContents := int64(999999999999)
							fewestContentsId := int32(0)
							contentMap := make(map[int32]int64)
							schematicID := int32(0)

							for _, c := range pin.Contents {
								contentMap[c.TypeId] = c.Amount
							}

							for srcPinID, route := range srcMap {
								srcPin := pinMap[srcPinID]
								if intArrayContains(aifTypes, srcPin.TypeId) && contentMap[route.ContentTypeId] < fewestContents {
									fewestContents = contentMap[route.ContentTypeId]
									fewestContentsId = route.ContentTypeId
									schematicID = srcPin.SchematicId
								}
							}

							if schematicID == 0 {
								continue
							}

							var schematicQuantity int64
							var err = h.clients.DB.QueryRow("select quantity from \"planetSchematicsTypeMap\" where \"isInput\" = true and \"typeID\" = $1 and \"schematicID\" = $2", fewestContentsId, schematicID).Scan(&schematicQuantity)
							if err != nil {
								logger.Error(err)
								return
							}

							var lastCycle time.Time
							for srcPinID, route := range srcMap {
								srcPin := pinMap[srcPinID]
								if route.ContentTypeId == fewestContentsId {
									ratePerHour = ratePerHour + schematicQuantity
									if lastCycle.Before(srcPin.LastCycleStart) {
										lastCycle = srcPin.LastCycleStart
									}
								}
							}

							if ratePerHour > 0 {
								hoursRemaining := int64(fewestContents / ratePerHour)
								lastCycle = lastCycle.Add(time.Hour * time.Duration(hoursRemaining+1))
								if lastCycle.Before(NextAttention) {
									NextAttention = lastCycle
								}
							}
						}
					}

					var planetName string
					var constelationName string

					err = h.clients.DB.QueryRow("select m1.\"itemName\" planet_name, (select m2.\"itemName\" constelation_name from \"mapDenormalize\" m2 where m2.\"itemID\" = m1.\"constellationID\") from \"mapDenormalize\" m1 where m1.\"itemID\" = $1", j.PlanetId).Scan(&planetName, &constelationName)
					if err != nil {
						logger.Error(err)
						return
					}

					bc <- PlanetData{
						//	Activity  string
						Account:          c.EVEAccountName.String,
						CharacterID:      c.ID,
						Character:        c.Name,
						PlanetName:       planetName,
						ConstelationName: constelationName,
						PlanetType:       j.PlanetType,
						BIFCount:         count,
						Extracted:        extracted,
						NextAttention:    NextAttention,
					}
				}(j)
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(bc)
	}()

	list := make([]PlanetData, 0)
	for b := range bc {
		list = append(list, b)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].NextAttention.Before(list[j].NextAttention)
	})

	return render("planets", w, list)
}
