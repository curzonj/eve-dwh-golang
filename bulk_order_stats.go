package main

import (
	"io/ioutil"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

func marketStatisticsPoller(d time.Duration) {
	logger := globals.logger.WithField("fn", "marketStatisticsPoller")
	logger.WithField("at", "start").Info()

	err := pollMarketStats()
	if err != nil {
		logger.Error(err)
	}

	for range time.Tick(d) {
		err := pollMarketStats()
		if err != nil {
			logger.Error(err)
		}
	}
}

func pollMarketStats() error {
	fetcher := &marketDataFetcher{}
	data, err := fetcher.GetOrderDataset(cfg.RegionID)
	if err != nil {
		return errors.Wrap(err, "fetching order data")
	}

	err = importBulkOrderStats(data)
	if err != nil {
		return errors.Wrap(err, "saving stats to pg")
	}

	return nil
}

func importBulkOrderStats(data orderDataset) error {
	dataTimestamp := time.Now().Unix()

	var storedTypeIDs []int32
	for _, id := range cfg.MarketGroups {
		var thisTypeIDs []int32
		err := globals.db.Select(&thisTypeIDs, "select \"typeID\" from \"invTypes\" where \"marketGroupID\" in (select market_group_id from market_group_arrays where id_list && '{"+strconv.Itoa(id)+"}')")
		if err != nil {
			return err
		}

		storedTypeIDs = append(storedTypeIDs, thisTypeIDs...)
	}

	insertSQL, err := ioutil.ReadFile("doc/bulk_order_stats_insert.sql")
	if err != nil {
		return errors.Wrap(err, "loading template sql")
	}

	tx, err := globals.db.Begin()
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

		_, err := stmt.Exec(typeID, cfg.RegionID, dataTimestamp, buyUnits, sellUnits)
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
