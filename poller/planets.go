package poller

import (
	"context"
	"fmt"
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
	Character *model.UserCharacter
	Planet    *esi.GetCharactersCharacterIdPlanets200Ok
	Error     error
}

type PlanetFetchResult struct {
	Character *model.UserCharacter
	Planet    *esi.GetCharactersCharacterIdPlanets200Ok
	Details   *esi.GetCharactersCharacterIdPlanetsPlanetIdOk
}

type PlanetData struct {
	CharacterID      int64
	Character        string
	PlanetName       string
	ConstelationName string
	PlanetType       string
	ProductName      string
	BIFCount         int
	Extracted        int32
	NextAttention    time.Time
	StorageFullAt    time.Time
}

func FetchPlanets(ctx context.Context, clients types.Clients, characters []*model.UserCharacter) (<-chan PlanetFetchResult, <-chan PlanetFetchError) {
	var wg sync.WaitGroup
	errorChan := make(chan PlanetFetchError)
	planetChan := make(chan PlanetFetchResult)

	for _, c := range characters {
		wg.Add(1)
		go func(c *model.UserCharacter) {
			defer wg.Done()

			ctx, err := c.TokenSourceContext(ctx, clients.ESIAuthenticator)
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
	characters, err := p.clients.DB.GetAllCharacters()
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
			calc := newPlanetCalculator(p.clients, result)
			obs, err := calc.buildPlanetObservation()
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

func BuildPlanetData(clients types.Clients, result PlanetFetchResult) (*PlanetData, error) {
	calc := newPlanetCalculator(clients, result)
	obs, err := calc.buildPlanetObservation()
	if err != nil {
		return nil, err
	}

	var planetName string
	err = clients.DB.QueryRow("select item_name from sde_names where item_id = $1", result.Planet.PlanetId).Scan(&planetName)
	if err != nil {
		return nil, errors.Wrap(err, "fetching planet name")
	}

	var constelationName string
	err = clients.DB.QueryRow("select item_name from sde_names where item_id = (select constellation_id from sde_solar_systems where $1 = ANY (planet_ids))", result.Planet.PlanetId).Scan(&constelationName)
	if err != nil {
		return nil, errors.Wrap(err, "fetching constelation name")
	}

	var P1ProductName string
	if obs.BasicFactoryOutputID != 0 {
		t, err := clients.DB.GetSDEType(obs.BasicFactoryOutputID)
		if err != nil {
			return nil, err
		}

		P1ProductName = t.Name
	}

	nextAttention, err := calc.nextAttention()
	if err != nil {
		return nil, errors.Wrap(err, "nextAttention")
	}

	storageFullAt, err := calc.storageFullAt()
	if err != nil {
		return nil, errors.Wrap(err, "storageFullAt")
	}

	return &PlanetData{
		CharacterID:      result.Character.ID,
		Character:        result.Character.Name,
		PlanetName:       planetName,
		ProductName:      P1ProductName,
		ConstelationName: constelationName,
		Extracted:        calc.QtyPerHour(),
		PlanetType:       result.Planet.PlanetType,
		NextAttention:    nextAttention,
		StorageFullAt:    storageFullAt,
	}, nil
}
