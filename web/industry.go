package web

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/antihax/goesi"
	"github.com/curzonj/eve-dwh-golang/model"
)

type BlueprintJob struct {
	//	Activity  string
	Installer string
	Type      string
	EndDate   time.Time
}

func (h *handler) industryJobs(w http.ResponseWriter, r *http.Request) error {
	session := session(r)
	userID := session.Values["user_id"].(string)
	logger := logger(r)

	var characters []model.UserCharacter
	err := h.clients.DB.Select(&characters, "select * from user_characters where user_id = $1", userID)
	if err != nil {
		return err
	}

	list := make([]BlueprintJob, 0)

	for _, c := range characters {
		tokSrc, err := c.TokenSource(h.clients)
		if err != nil {
			logger.Error(err)
			continue
		}

		ctx := context.WithValue(r.Context(), goesi.ContextOAuth2, tokSrc)
		data, _, err := h.clients.ESIClient.IndustryApi.GetCharactersCharacterIdIndustryJobs(ctx, int32(c.ID), map[string]interface{}{
			"includeCompleted": false,
		})

		if err != nil {
			logger.Error(err)
			continue
		}

		for _, j := range data {
			var name string
			err := h.clients.DB.Get(&name, "select \"typeName\" from \"invTypes\" where \"typeID\" = $1 limit 1", j.BlueprintTypeId)
			if err != nil {
				logger.Error(err)
				continue
			}

			list = append(list, BlueprintJob{
				Installer: c.Name,
				Type:      name,
				EndDate:   j.EndDate,
			})
		}

	}

	w.Write([]byte(fmt.Sprintf("welcome %s: %+v", userID, list)))
	return nil
}
