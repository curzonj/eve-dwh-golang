package web

import (
	"net/http"
	"sort"
	"sync"

	"github.com/antihax/goesi/esi"
	"github.com/curzonj/eve-dwh-golang/model"
)

type PlanetData struct {
	//	Activity  string
	CharacterID int64
	Character   string
	PlanetID    int32
	PlanetType  string
	BIFCount    int
}

var (
	aifTypes       = []int32{2470, 2472, 2474, 2480, 2484, 2485, 2491, 2494}
	launchpadTypes = []int32{2256, 2542, 2543, 2544, 2552, 2555, 2556, 2557}
	storageTypes   = []int32{2257, 2535, 2536, 2541, 2558, 2560, 2561, 2562}
	bifTypes       = []int32{2469, 2471, 2473, 2481, 2483, 2490, 2492, 2493}
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

					for _, pin := range data.Pins {
						if intArrayContains(bifTypes, pin.TypeId) {
							count = count + 1
						}
					}

					bc <- PlanetData{
						//	Activity  string
						CharacterID: c.ID,
						Character:   c.Name,
						PlanetID:    j.PlanetId,
						PlanetType:  j.PlanetType,
						BIFCount:    count,
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
		if list[i].CharacterID < list[j].CharacterID {
			return true
		}
		if list[i].CharacterID > list[j].CharacterID {
			return false
		}
		return list[i].PlanetID < list[j].PlanetID
	})

	return render("planets", w, list)
}
