package main

import (
	"io/ioutil"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
	cli "gopkg.in/urfave/cli.v1"
)

func importAction(c *cli.Context) error {
	connectToDatabase()

	fetcher := &marketDataFetcher{}
	data, err := fetcher.GetOrderDataset(cfg.RegionID)
	if err != nil {
		return errors.Wrap(err, "fetching order data")
	}

	return importBulkOrderStats(data)
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
		globals.logger.WithFields(log.Fields{
			"fn":     "importBulkOrderStats",
			"typeID": typeID,
			"orders": len(orders),
		}).Info()

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
