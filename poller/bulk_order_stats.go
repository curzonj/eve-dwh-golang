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
	p := &pollerHandler{
		clients: clients,
		logger:  clients.Logger.WithField("fn", "marketStatisticsPoller"),
		cfg:     cfg,
	}

	p.leadingEdgeTick(cfg.Interval, p.pollMarketStats)
}

func (p *pollerHandler) pollMarketStats() error {
	for _, regionID := range whitelistRegions {
		err := p.pollRegion(regionID)
		if err != nil {
			p.logger.Error(err)
		}
	}

	return nil
}

func (p *pollerHandler) pollRegion(regionID int32) error {
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

func (p *pollerHandler) importBulkOrderStats(regionID int32, data orderDataset) error {
	dataTimestamp := time.Now().Unix()

	var blacklistTypeIDs []int32
	for _, id := range blacklistGroups {
		var thisTypeIDs []int32
		err := p.clients.DB.Select(&thisTypeIDs, "select type_id from sde_types where market_group_id in (select sde_market_group_arrays.market_group_id from sde_market_group_arrays where id_list && '{"+strconv.Itoa(int(id))+"}')")
		if err != nil {
			return err
		}

		blacklistTypeIDs = append(blacklistTypeIDs, thisTypeIDs...)
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
		if intArrayContains(blacklistTypeIDs, typeID) {
			continue
		}

		var buyUnits, sellUnits int64
		var buyOrderCount, sellOrderCount int64
		var buyPriceMaxIsk, sellPriceMinIsk int64

		for _, o := range orders {
			iskPrice := int64(o.Price * 100)

			if o.IsBuyOrder {
				buyUnits = buyUnits + int64(o.VolumeRemain)
				buyOrderCount = buyOrderCount + 1
				if iskPrice > buyPriceMaxIsk {
					buyPriceMaxIsk = iskPrice
				}
			} else {
				sellUnits = sellUnits + int64(o.VolumeRemain)
				sellOrderCount = sellOrderCount + 1
				if sellPriceMinIsk == 0 || iskPrice < sellPriceMinIsk {
					sellPriceMinIsk = iskPrice
				}
			}
		}

		_, err := stmt.Exec(
			typeID, regionID, dataTimestamp,
			buyUnits, sellUnits,
			buyPriceMaxIsk, sellPriceMinIsk,
			buyOrderCount, sellOrderCount)
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
