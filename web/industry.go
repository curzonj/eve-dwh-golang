package web

import (
	"net/http"
	"sort"
	"sync"
	"time"

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
	logger := logger(r.Context())

	characters, err := h.clients.DB.GetUserCharactersByUserID(userID)
	if err != nil {
		return err
	}

	bc := make(chan BlueprintJob)
	var wg sync.WaitGroup

	for _, c := range characters {
		wg.Add(1)
		go func(c *model.UserCharacter) {
			defer wg.Done()

			ctx, err := c.TokenSourceContext(r.Context(), h.clients.ESIAuthenticator)
			if err != nil {
				logger.Error(err)
				return
			}

			data, _, err := h.clients.EVEBreakerClient.ESI.IndustryApi.GetCharactersCharacterIdIndustryJobs(ctx, int32(c.ID), map[string]interface{}{
				"includeCompleted": false,
			})

			if err != nil {
				logger.Error(err)
				return
			}

			for _, j := range data {
				t, err := h.clients.DB.GetSDEType(j.BlueprintTypeId)
				if err != nil {
					logger.Error(err)
					return
				}

				bc <- BlueprintJob{
					Installer: c.Name,
					Type:      t.Name,
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
