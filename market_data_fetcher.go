package main

import (
	"strconv"

	"github.com/antihax/goesi/esi"
)

type orderDataset map[int32][]esi.GetMarketsRegionIdOrders200Ok

type marketDataFetcher struct{}

func (f *marketDataFetcher) GetOrderDataset(regionID int32) (orderDataset, error) {
	dataset := make(orderDataset)

	eachPage, err := f.fetchOrders(regionID)
	if err != nil {
		return nil, err
	}

	for _, page := range eachPage {
		for _, order := range page {
			original := dataset[order.TypeId]
			dataset[order.TypeId] = append(original, order)
		}
	}

	return dataset, nil
}

func (f *marketDataFetcher) fetchOrders(regionID int32) ([][]esi.GetMarketsRegionIdOrders200Ok, error) {
	firstPage, resp, err := globals.esiClient.MarketApi.GetMarketsRegionIdOrders("buy", regionID, nil)
	if err != nil {
		return nil, err
	}

	pages, err := strconv.Atoi(resp.Header.Get("X-Pages"))
	if err != nil {
		return nil, err
	}

	allPages, err := f.remainingOrderPages(regionID, pages)
	if err != nil {
		return nil, err
	}

	allPages = append(allPages, firstPage)

	return allPages, nil
}

func (f *marketDataFetcher) remainingOrderPages(regionID int32, lastPage int) ([][]esi.GetMarketsRegionIdOrders200Ok, error) {
	pages := lastPage - 1
	results := make(chan []esi.GetMarketsRegionIdOrders200Ok, pages)
	errors := make(chan error, pages)

	for i := 2; i <= lastPage; i++ {
		go func(page int32) {
			params := make(map[string]interface{})
			params["page"] = page

			list, _, err := globals.esiClient.MarketApi.GetMarketsRegionIdOrders("buy", regionID, params)

			if err != nil {
				errors <- err
			} else {
				results <- list
			}
		}(int32(i))
	}

	list := make([][]esi.GetMarketsRegionIdOrders200Ok, 0, pages)

	for {
		select {
		case r := <-results:
			list = append(list, r)
			if len(list) == pages {
				return list, nil
			}
		case err := <-errors:
			return nil, err
		}
	}
}
