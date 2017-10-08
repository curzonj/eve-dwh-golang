package web

import (
	"context"
	"net/http"
	"sort"
	"sync"
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

	bc := make(chan BlueprintJob)
	var wg sync.WaitGroup

	for _, c := range characters {
		wg.Add(1)
		go func(c model.UserCharacter) {
			defer wg.Done()

			tokSrc, err := c.TokenSource(h.clients)
			if err != nil {
				logger.Error(err)
				return
			}

			ctx := context.WithValue(r.Context(), goesi.ContextOAuth2, tokSrc)
			data, _, err := h.clients.ESIClient.IndustryApi.GetCharactersCharacterIdIndustryJobs(ctx, int32(c.ID), map[string]interface{}{
				"includeCompleted": false,
			})

			if err != nil {
				logger.Error(err)
				return
			}

			for _, j := range data {
				var name string
				err := h.clients.DB.Get(&name, "select \"typeName\" from \"invTypes\" where \"typeID\" = $1 limit 1", j.BlueprintTypeId)
				if err != nil {
					logger.Error(err)
					return
				}

				bc <- BlueprintJob{
					Installer: c.Name,
					Type:      name,
					EndDate:   j.EndDate,
				}
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(bc)
	}()

	list := make([]BlueprintJob, 0)
	for b := range bc {
		list = append(list, b)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[j].EndDate.After(list[i].EndDate)
	})

	return render("industry", w, list)
}
