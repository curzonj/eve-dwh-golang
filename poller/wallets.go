package poller

import (
	"context"
	"fmt"
	"time"

	"github.com/curzonj/eve-dwh-golang/model"
	"github.com/curzonj/eve-dwh-golang/types"
	"github.com/curzonj/eve-dwh-golang/utils/sqlh"
)

func WalletsPoller(clients types.Clients) {
	p := &pollerHandler{
		clients: clients,
		logger:  clients.Logger.WithField("fn", "walletsPoller"),
	}

	p.leadingEdgeTick(time.Hour, p.walletsPollerTick)
}

func (p *pollerHandler) getCharacterESIWalletTransactions(ctx context.Context, c model.UserCharacter) ([]map[string]interface{}, error) {
	transactions, _, err := p.clients.EVERetryClient.ESI.WalletApi.GetCharactersCharacterIdWalletTransactions(ctx, int32(c.ID), nil)
	if err != nil {
		return nil, err
	}

	list := make([]map[string]interface{}, 0, len(transactions))

	for _, row := range transactions {
		list = append(list, map[string]interface{}{
			"entity_id":      c.ID,
			"division_id":    0,
			"transaction_id": row.TransactionId,
			"character_id":   c.ID,
			"occured_at":     row.Date,
			"quantity":       row.Quantity,
			"type_id":        row.TypeId,
			"price":          int64(row.UnitPrice * 100),
			"buy":            row.IsBuy,
			"is_personal":    row.IsPersonal,
			"journal_ref_id": row.JournalRefId,
			"client_id":      row.ClientId,
			"location_id":    row.LocationId,
		})
	}

	return list, nil
}

func (p *pollerHandler) getCharacterESIWalletJournal(ctx context.Context, c model.UserCharacter) ([]map[string]interface{}, error) {
	journals, _, err := p.clients.EVERetryClient.ESI.WalletApi.GetCharactersCharacterIdWalletJournal(ctx, int32(c.ID), nil)
	if err != nil {
		return nil, err
	}

	list := make([]map[string]interface{}, 0, len(journals))

	for _, row := range journals {
		extra_info, _ := row.ExtraInfo.MarshalJSON()
		list = append(list, map[string]interface{}{
			"entity_id":        c.ID,
			"division_id":      0,
			"journal_ref_id":   row.RefId,
			"ref_type":         row.RefType,
			"occured_at":       row.Date,
			"reason":           row.Reason,
			"party_1_id":       row.FirstPartyId,
			"party_1_type":     row.FirstPartyType,
			"party_2_id":       row.SecondPartyId,
			"party_2_type":     row.SecondPartyType,
			"amount":           int64(row.Amount * 100),
			"balance":          int64(row.Balance * 100),
			"tax_collector_id": row.TaxRecieverId,
			"tax_amount":       int64(row.Tax * 100),
			"extra_info":       extra_info,
		})
	}

	return list, nil
}

func (p *pollerHandler) walletsPollerTick() error {
	var characters []model.UserCharacter
	err := p.clients.DB.Select(&characters, "select * from user_characters")
	if err != nil {
		return err
	}

	for _, c := range characters {
		l := p.logger.WithField("characterName", c.Name)
		ctx := context.WithValue(context.TODO(), types.ContextLoggerKey, l)

		ctx, err := c.TokenSourceContext(ctx, p.clients)
		if err != nil {
			l.Error(err)
			continue
		}

		journals, err := p.getCharacterESIWalletJournal(ctx, c)
		if err != nil {
			l.Error(err)
			continue
		}

		for _, values := range journals {
			columns := sqlh.BuildColumnsValues(values)
			_, err = p.clients.DB.NamedExec(fmt.Sprintf("INSERT INTO wallet_journals %s ON CONFLICT DO NOTHING", columns), values)
			if err != nil {
				l.Error(err)
			}
		}

		transactions, err := p.getCharacterESIWalletTransactions(ctx, c)
		if err != nil {
			l.Error(err)
			continue
		}

		for _, values := range transactions {
			columns := sqlh.BuildColumnsValues(values)
			_, err = p.clients.DB.NamedExec(fmt.Sprintf("INSERT INTO wallet_transactions %s ON CONFLICT DO NOTHING", columns), values)
			if err != nil {
				l.Error(err)
			}
		}
	}

	return nil
}
