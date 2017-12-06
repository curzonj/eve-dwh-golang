package poller

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/antihax/goesi/esi"
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

func (p *pollerHandler) cacheOrderDataset(regionID int32, data orderDataset) error {
	var buf bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&buf) // Will write to network.

	err := enc.Encode(data)
	if err != nil {
		return errors.Wrap(err, "gobEncode")
	}

	key := fmt.Sprintf("poller:market_data:region:%d", regionID)

	return p.clients.DB.SetCacheItem(key, buf.Bytes())
}

func (p *pollerHandler) getCachedOrderDataset(regionID int32) (orderDataset, error) {
	key := fmt.Sprintf("poller:market_data:region:%d", regionID)
	byteA, err := p.clients.DB.GetCacheItem(key)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(byteA)
	dec := gob.NewDecoder(buf) // Will write to network.

	var data orderDataset
	err = dec.Decode(&data)
	if err != nil {
		return nil, errors.Wrap(err, "gobDecode")
	}

	return data, nil
}

func (p *pollerHandler) pollRegion(regionID int32) error {
	fetcher := &marketDataFetcher{
		clients: p.clients,
	}

	previous, err := p.getCachedOrderDataset(regionID)
	if err != nil {
		p.logger.Error(errors.Wrap(err, "getCachedOrderDataset"))
		previous = make(orderDataset)
	}

	data, err := fetcher.GetOrderDataset(regionID)
	if err != nil {
		return errors.Wrap(err, "GetOrderDataset")
	}

	err = p.cacheOrderDataset(regionID, data)
	if err != nil {
		return errors.Wrap(err, "cacheOrderDataset")
	}

	err = p.importBulkOrderStats(regionID, data, previous)
	if err != nil {
		return errors.Wrap(err, "importBulkOrderStats")
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

func (p *pollerHandler) importBulkOrderStats(regionID int32, data orderDataset, previous orderDataset) error {
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

		s := generateIdentityOrderStats(orders)

		if po, ok := previous[typeID]; ok {
			ps := generateIdentityOrderStats(po)

			if s == ps {
				continue
			}
		}

		_, err := stmt.Exec(
			typeID, regionID, dataTimestamp,
			s.buyUnits, s.sellUnits,
			s.buyPriceMaxIsk, s.sellPriceMinIsk,
			s.buyOrderCount, s.sellOrderCount)
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

type identityOrderStats struct {
	buyUnits        int64
	sellUnits       int64
	buyOrderCount   int64
	sellOrderCount  int64
	buyPriceMaxIsk  int64
	sellPriceMinIsk int64
}

/*
 * This will differ from other stats that will be delta stats. These stats
 * will be ignored if they match the previous stats. Delta stats will be
 * ignored if they are zero.
 */
func generateIdentityOrderStats(orders []esi.GetMarketsRegionIdOrders200Ok) identityOrderStats {
	var s identityOrderStats

	for _, o := range orders {
		iskPrice := int64(o.Price * 100)

		if o.IsBuyOrder {
			s.buyUnits = s.buyUnits + int64(o.VolumeRemain)
			s.buyOrderCount = s.buyOrderCount + 1
			if iskPrice > s.buyPriceMaxIsk {
				s.buyPriceMaxIsk = iskPrice
			}
		} else {
			s.sellUnits = s.sellUnits + int64(o.VolumeRemain)
			s.sellOrderCount = s.sellOrderCount + 1
			if s.sellPriceMinIsk == 0 || iskPrice < s.sellPriceMinIsk {
				s.sellPriceMinIsk = iskPrice
			}
		}
	}

	return s
}
