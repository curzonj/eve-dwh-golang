package web

import (
	"context"
	"net/http"
	"sort"

	log "github.com/Sirupsen/logrus"

	"github.com/curzonj/eve-dwh-golang/model"
	"github.com/curzonj/eve-dwh-golang/poller"
)

func (h *handler) buildPlanetList(ctx context.Context, characters []*model.UserCharacter) []*poller.PlanetData {
	logger := logger(ctx)
	list := make([]*poller.PlanetData, 0)
	planetC, errC := poller.FetchPlanets(ctx, h.clients, characters)

	for {
		select {
		case result, ok := <-errC:
			if !ok {
				return list
			}

			l := logger.WithField("character_id", result.Character.ID)

			if result.Planet != nil {
				l = logger.WithField("planet_id", result.Planet.PlanetId)
			}

			l.Error(result.Error)
		case result := <-planetC:
			data, err := poller.BuildPlanetData(h.clients, result)
			if err != nil {
				logger.WithFields(log.Fields{
					"planet_id":    result.Planet.PlanetId,
					"character_id": result.Character.ID,
				}).Error(err)

				continue
			}

			list = append(list, data)
		}
	}
}

func (h *handler) planets(w http.ResponseWriter, r *http.Request) error {
	session := session(r)
	userID := session.Values["user_id"].(string)

	characters, err := h.clients.DB.GetUserCharactersByUserID(userID)
	if err != nil {
		return err
	}

	list := h.buildPlanetList(r.Context(), characters)
	sort.Slice(list, func(i, j int) bool {
		return list[i].NextAttention.Before(list[j].NextAttention)
	})

	return render("planets", w, list)
}
