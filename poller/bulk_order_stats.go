package poller

import (
	"io/ioutil"
	"strconv"
	"time"

	"github.com/curzonj/eve-dwh-golang/types"
	"github.com/pkg/errors"
)

var blacklistGroups = []int32{2, 150, 1396, 350001, 1659, 1954}

// Domain The Forge Heimatar Metropolis Sinq Laison
var whitelistRegions = []int32{10000043, 10000002, 10000030, 10000042, 10000032}

func MarketStatisticsPoller(clients types.Clients, cfg Cfg) {
	p := &poller{
		clients: clients,
		logger:  clients.Logger.WithField("fn", "marketStatisticsPoller"),
		cfg:     cfg,
	}

	p.leadingEdgeTick(cfg.Interval, p.pollMarketStats)
}

func (p *poller) pollMarketStats() error {
	for _, regionID := range whitelistRegions {
		err := p.pollRegion(regionID)
		if err != nil {
			p.logger.Error(err)
		}
	}

	return nil
}

func (p *poller) pollRegion(regionID int32) error {
	fetcher := &marketDataFetcher{
		clients: p.clients,
	}

	data, err := fetcher.GetOrderDataset(regionID)
	if err != nil {
		return errors.Wrap(err, "fetching order data")
	}

	err = p.importBulkOrderStats(regionID, data)
	if err != nil {
		return errors.Wrap(err, "saving stats to pg")
	}

	return nil
}

func intArrayContains(list []int32, a int32) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (p *poller) importBulkOrderStats(regionID int32, data orderDataset) error {
	dataTimestamp := time.Now().Unix()

	var storedTypeIDs []int32
	for _, id := range blacklistGroups {
		var thisTypeIDs []int32
		err := p.clients.DB.Select(&thisTypeIDs, "select type_id from sde_types where market_group_id in (select sde_market_group_arrays.market_group_id from sde_market_group_arrays where id_list && '{"+strconv.Itoa(int(id))+"}')")
		if err != nil {
			return err
		}

		storedTypeIDs = append(storedTypeIDs, thisTypeIDs...)
	}

	insertSQL, err := ioutil.ReadFile("doc/bulk_order_stats_insert.sql")
	if err != nil {
		return errors.Wrap(err, "loading template sql")
	}

	tx, err := p.clients.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(string(insertSQL))
	if err != nil {
		return errors.Wrap(err, "preparing insert sql")
	}
	defer stmt.Close()

	for typeID, orders := range data {
		if intArrayContains(blacklistGroups, typeID) {
			continue
		}

		var buyUnits, sellUnits int64

		for _, o := range orders {
			if o.IsBuyOrder {
				buyUnits = buyUnits + int64(o.VolumeRemain)
			} else {
				sellUnits = sellUnits + int64(o.VolumeRemain)
			}
		}

		_, err := stmt.Exec(typeID, regionID, dataTimestamp, buyUnits, sellUnits)
		if err != nil {
			return err
		}
	}

	tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
