package poller

import (
	"io/ioutil"
	"strconv"
	"time"

	"github.com/curzonj/eve-dwh-golang/types"
	"github.com/pkg/errors"
)

func MarketStatisticsPoller(clients types.Clients, cfg Cfg) {
	logger := clients.Logger.WithField("fn", "marketStatisticsPoller")
	logger.WithField("at", "start").Info()

	p := &poller{
		clients: clients,
		cfg:     cfg,
	}

	err := p.pollMarketStats()
	if err != nil {
		logger.Error(err)
	}

	for range time.Tick(cfg.Interval) {
		err := p.pollMarketStats()
		if err != nil {
			logger.Error(err)
		}
	}
}

type poller struct {
	clients types.Clients
	cfg     Cfg
}

func (p *poller) pollMarketStats() error {
	fetcher := &marketDataFetcher{
		clients: p.clients,
	}

	data, err := fetcher.GetOrderDataset(p.cfg.RegionID)
	if err != nil {
		return errors.Wrap(err, "fetching order data")
	}

	err = p.importBulkOrderStats(data)
	if err != nil {
		return errors.Wrap(err, "saving stats to pg")
	}

	return nil
}

func (p *poller) importBulkOrderStats(data orderDataset) error {
	dataTimestamp := time.Now().Unix()

	var storedTypeIDs []int32
	for _, id := range p.cfg.MarketGroups {
		var thisTypeIDs []int32
		err := p.clients.DB.Select(&thisTypeIDs, "select \"typeID\" from \"invTypes\" where \"marketGroupID\" in (select market_group_id from market_group_arrays where id_list && '{"+strconv.Itoa(id)+"}')")
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

	for _, typeID := range storedTypeIDs {
		orders := data[typeID]
		var buyUnits, sellUnits int64

		for _, o := range orders {
			if o.IsBuyOrder {
				buyUnits = buyUnits + int64(o.VolumeRemain)
			} else {
				sellUnits = sellUnits + int64(o.VolumeRemain)
			}
		}

		_, err := stmt.Exec(typeID, p.cfg.RegionID, dataTimestamp, buyUnits, sellUnits)
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
