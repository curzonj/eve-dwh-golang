package main

import (
	"strconv"
	"sync"

	"github.com/antihax/goesi/esi"
	log "github.com/sirupsen/logrus"
)

type orderDataset map[int32][]esi.GetMarketsRegionIdOrders200Ok

type marketDataFetcher struct{}

func (f *marketDataFetcher) GetOrderDataset(regionID int32) (orderDataset, error) {
	dataset := make(orderDataset)

	eachPage, err := f.FetchOrders(regionID)
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

func (f *marketDataFetcher) FetchOrders(regionID int32) ([][]esi.GetMarketsRegionIdOrders200Ok, error) {
	firstPage, resp, err := globals.esiClient.MarketApi.GetMarketsRegionIdOrders("buy", regionID, nil)
	log.WithFields(log.Fields{
		"page":  1,
		"error": err,
		"count": len(firstPage),
	}).Info()

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

func (f *marketDataFetcher) remainingOrderPages(regionID int32, pages int) ([][]esi.GetMarketsRegionIdOrders200Ok, error) {
	var wg sync.WaitGroup
	results := make(chan []esi.GetMarketsRegionIdOrders200Ok, pages)
	errors := make(chan error, pages)

	for i := 2; i <= pages; i++ {
		wg.Add(1)
		go f.fetchPage(&wg, results, errors, regionID, int32(i))
	}

	wg.Wait()
	close(results)

	list := make([][]esi.GetMarketsRegionIdOrders200Ok, 0, pages-1)
	for r := range results {
		list = append(list, r)
	}

	return list, nil
}

func (f *marketDataFetcher) fetchPage(wg *sync.WaitGroup, results chan<- []esi.GetMarketsRegionIdOrders200Ok, errors chan<- error, regionID, page int32) {
	defer wg.Done()
	params := make(map[string]interface{})
	params["page"] = page
	list, _, err := globals.esiClient.MarketApi.GetMarketsRegionIdOrders("buy", regionID, params)
	log.WithFields(log.Fields{
		"page":  page,
		"error": err,
		"count": len(list),
	}).Info()

	if err != nil {
		errors <- err
	} else {
		results <- list
	}
}
