package poller

import (
	"context"
	"fmt"
	"math"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/antihax/goesi/esi"
	"github.com/curzonj/eve-dwh-golang/model"
	"github.com/curzonj/eve-dwh-golang/types"
	"github.com/curzonj/eve-dwh-golang/utils/sqlh"
	"github.com/pkg/errors"
)

type PlanetFetchError struct {
	Character model.UserCharacter
	Planet    *esi.GetCharactersCharacterIdPlanets200Ok
	Error     error
}

type PlanetFetchResult struct {
	Character model.UserCharacter
	Planet    *esi.GetCharactersCharacterIdPlanets200Ok
	Details   *esi.GetCharactersCharacterIdPlanetsPlanetIdOk
}

type planetObservation struct {
	ExtractionTypeID int32
	QtyPerCycle      int32
	CycleTime        int32
	ExtractorHeads   int32
	Extractors       int32
	HeadRadius       float32
	BasicFactories   int32
	UpgradeLevel     int32
}

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

func FetchPlanets(ctx context.Context, clients types.Clients, characters []model.UserCharacter) (<-chan PlanetFetchResult, <-chan PlanetFetchError) {
	var wg sync.WaitGroup
	errorChan := make(chan PlanetFetchError)
	planetChan := make(chan PlanetFetchResult)

	for _, c := range characters {
		wg.Add(1)
		go func(c model.UserCharacter) {
			defer wg.Done()

			ctx, err := c.TokenSourceContext(ctx, clients)
			if err != nil {
				errorChan <- PlanetFetchError{c, nil, err}
				return
			}

			data, _, err := clients.EVEBreakerClient.ESI.PlanetaryInteractionApi.GetCharactersCharacterIdPlanets(ctx, int32(c.ID), nil)
			if err != nil {
				errorChan <- PlanetFetchError{c, nil, err}
				return
			}

			for _, j := range data {
				wg.Add(1)
				go func(j esi.GetCharactersCharacterIdPlanets200Ok) {
					defer wg.Done()

					data, _, err := clients.EVEBreakerClient.ESI.PlanetaryInteractionApi.GetCharactersCharacterIdPlanetsPlanetId(ctx, int32(c.ID), j.PlanetId, nil)
					if err != nil {
						errorChan <- PlanetFetchError{c, &j, err}
						return
					}

					planetChan <- PlanetFetchResult{c, &j, &data}
				}(j)
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(errorChan)
	}()

	return planetChan, errorChan
}

func PlanetsPoller(clients types.Clients) {
	p := &pollerHandler{
		clients: clients,
		logger:  clients.Logger.WithField("fn", "planetsPoller"),
	}

	p.leadingEdgeTick(time.Duration(24)*time.Hour, p.planetsPollerTick)
}

func (p *pollerHandler) planetsPollerTick() error {
	var characters []model.UserCharacter
	err := p.clients.DB.Select(&characters, "select * from user_characters")
	if err != nil {
		return err
	}

	planetC, errC := FetchPlanets(context.TODO(), p.clients, characters)

	for {
		select {
		case result, ok := <-errC:
			if !ok {
				return nil
			}

			p.logger.Error(result.Error)
		case result := <-planetC:
			obs, err := buildPlanetObservation(result)
			if err != nil {
				p.logger.WithFields(log.Fields{
					"planet_id":    result.Planet.PlanetId,
					"character_id": result.Character.ID,
				}).Error(err)
				continue
			}

			if obs.Extractors == 0 {
				continue
			}

			values := map[string]interface{}{
				"character_id":       result.Character.ID,
				"planet_id":          result.Planet.PlanetId,
				"observed_at":        time.Now(),
				"extraction_type_id": obs.ExtractionTypeID,
				"qty_per_cycle":      obs.QtyPerCycle,
				"cycle_time":         obs.CycleTime,
				"extractor_heads":    obs.ExtractorHeads,
				"extractors":         obs.Extractors,
				"head_radius":        obs.HeadRadius,
				"basic_factories":    obs.BasicFactories,
				"upgrade_level":      obs.UpgradeLevel,
			}

			columns := sqlh.BuildColumnsValues(values)
			_, err = p.clients.DB.NamedExec(fmt.Sprintf("INSERT INTO planet_extraction_history %s ON CONFLICT DO NOTHING", columns), values)
			if err != nil {
				p.logger.Error(err)
			}
		}
	}
}

func buildPlanetObservation(result PlanetFetchResult) (*planetObservation, error) {
	data := result.Details
	obs := &planetObservation{
		UpgradeLevel: result.Planet.UpgradeLevel,
	}

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
			obs.BasicFactories = obs.BasicFactories + 1
		}

		if intArrayContains(extractorTypes, pin.TypeId) {
			obs.Extractors = obs.Extractors + 1
			obs.QtyPerCycle = obs.QtyPerCycle + pin.ExtractorDetails.QtyPerCycle
			obs.ExtractorHeads = obs.ExtractorHeads + int32(len(pin.ExtractorDetails.Heads))

			if obs.CycleTime != pin.ExtractorDetails.CycleTime {
				if obs.CycleTime != 0 {
					return nil, errors.Errorf("inconsistent extractor cycle times, %d != %d", obs.CycleTime, pin.ExtractorDetails.CycleTime)
				}

				obs.CycleTime = pin.ExtractorDetails.CycleTime
			}

			if obs.ExtractionTypeID != pin.ExtractorDetails.ProductTypeId {
				if obs.ExtractionTypeID != 0 {
					return nil, errors.Errorf("inconsistent extraction types, %d != %d", obs.ExtractionTypeID, pin.ExtractorDetails.ProductTypeId)
				}

				obs.ExtractionTypeID = pin.ExtractorDetails.ProductTypeId
			}

			if obs.HeadRadius != pin.ExtractorDetails.HeadRadius {
				if obs.HeadRadius != 0 {
					if math.Abs(float64(obs.HeadRadius-pin.ExtractorDetails.HeadRadius)) > 0.00006 {
						return nil, errors.Errorf("inconsistent head radius, %f != %f", obs.HeadRadius, pin.ExtractorDetails.HeadRadius)
					}
				} else {
					obs.HeadRadius = pin.ExtractorDetails.HeadRadius
				}
			}
		}
	}

	return obs, nil
}

func BuildPlanetData(clients types.Clients, result PlanetFetchResult) (*PlanetData, error) {
	data := result.Details
	count := 0
	extracted := 0
	nextAttention := time.Now().Add(time.Hour * time.Duration(10000))
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
			if pin.ExpiryTime.Before(nextAttention) {
				nextAttention = pin.ExpiryTime
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
			err := clients.DB.QueryRow("select contents->'inputs'->$1 from sde_planetary_schematics where schematic_id = $2", fewestContentsId, schematicID).Scan(&schematicQuantity)
			if err != nil {
				return nil, errors.Wrap(err, "fetching schematics")
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
				if lastCycle.Before(nextAttention) {
					nextAttention = lastCycle
				}
			}
		}
	}

	var planetName string
	err := clients.DB.QueryRow("select item_name from sde_names where item_id = $1", result.Planet.PlanetId).Scan(&planetName)
	if err != nil {
		return nil, errors.Wrap(err, "fetching planet name")
	}

	var constelationName string
	err = clients.DB.QueryRow("select item_name from sde_names where item_id = (select constellation_id from sde_solar_systems where $1 = ANY (planet_ids))", result.Planet.PlanetId).Scan(&constelationName)
	if err != nil {
		return nil, errors.Wrap(err, "fetching constelation name")
	}

	return &PlanetData{
		Account:          result.Character.EVEAccountName.String,
		CharacterID:      result.Character.ID,
		Character:        result.Character.Name,
		PlanetName:       planetName,
		ConstelationName: constelationName,
		PlanetType:       result.Planet.PlanetType,
		BIFCount:         count,
		Extracted:        extracted,
		NextAttention:    nextAttention,
	}, nil
}
