package poller

import (
	"math"
	"time"

	"github.com/antihax/goesi/esi"
	"github.com/curzonj/eve-dwh-golang/types"
	"github.com/pkg/errors"
)

var (
	aifTypes       = []int32{2470, 2472, 2474, 2480, 2484, 2485, 2491, 2494}
	launchpadTypes = []int32{2256, 2542, 2543, 2544, 2552, 2555, 2556, 2557}
	storageTypes   = []int32{2257, 2535, 2536, 2541, 2558, 2560, 2561, 2562}
	bifTypes       = []int32{2469, 2471, 2473, 2481, 2483, 2490, 2492, 2493}
	extractorTypes = []int32{2848, 3060, 3061, 3062, 3063, 3064, 3067, 3068}
)

type planetCalculator struct {
	clients        types.Clients
	pinMap         map[int64]esi.GetCharactersCharacterIdPlanetsPlanetIdPin
	sourceRouteMap map[int64]map[int64]esi.GetCharactersCharacterIdPlanetsPlanetIdRoute
	destRouteMap   map[int64]map[int64]esi.GetCharactersCharacterIdPlanetsPlanetIdRoute
	result         PlanetFetchResult
}

type planetObservation struct {
	ExtractionTypeID     int32
	QtyPerCycle          int32
	CycleTime            int32
	ExtractorHeads       int32
	Extractors           int32
	HeadRadius           float32
	BasicFactories       int32
	BasicFactoryOutputID int32
	UpgradeLevel         int32
}

func newPlanetCalculator(clients types.Clients, result PlanetFetchResult) *planetCalculator {
	data := result.Details

	calc := &planetCalculator{
		clients:        clients,
		result:         result,
		pinMap:         make(map[int64]esi.GetCharactersCharacterIdPlanetsPlanetIdPin, len(data.Pins)),
		sourceRouteMap: make(map[int64]map[int64]esi.GetCharactersCharacterIdPlanetsPlanetIdRoute),
		destRouteMap:   make(map[int64]map[int64]esi.GetCharactersCharacterIdPlanetsPlanetIdRoute),
	}

	for _, r := range data.Routes {
		if calc.sourceRouteMap[r.SourcePinId] == nil {
			calc.sourceRouteMap[r.SourcePinId] = make(map[int64]esi.GetCharactersCharacterIdPlanetsPlanetIdRoute)
		}

		if calc.destRouteMap[r.DestinationPinId] == nil {
			calc.destRouteMap[r.DestinationPinId] = make(map[int64]esi.GetCharactersCharacterIdPlanetsPlanetIdRoute)
		}

		calc.sourceRouteMap[r.SourcePinId][r.DestinationPinId] = r
		calc.destRouteMap[r.DestinationPinId][r.SourcePinId] = r
	}

	for _, pin := range data.Pins {
		calc.pinMap[pin.PinId] = pin
	}

	return calc
}

func (c *planetCalculator) countBIFs() int {
	count := 0

	for _, pin := range c.result.Details.Pins {
		if intArrayContains(bifTypes, pin.TypeId) {
			count = count + 1
		}
	}

	return count
}

func (c *planetCalculator) storageFullAt() (time.Time, error) {
	var zeroTime time.Time
	earliestFullAt := time.Now().Add(time.Hour * time.Duration(10000)).Truncate(time.Minute)

	if len(c.result.Details.Pins) == 1 {
		return zeroTime, nil
	}

	for _, pin := range c.result.Details.Pins {
		if !intArrayContains(launchpadTypes, pin.TypeId) {
			continue
		}

		srcMap := c.destRouteMap[pin.PinId]
		capacity := 10000.0
		currentContentVolume := float64(0)
		fillRate := float64(0)

		// If it has an extractor feeding to it, check the half cycle time based on 3/4 of the total output
		// For all LPs, check the ratePer hour and calculate number of hours from the earliest last cycle

		for _, contents := range pin.Contents {
			t, err := c.clients.DB.GetSDEType(contents.TypeId)
			if err != nil {
				return zeroTime, err
			}

			currentContentVolume = currentContentVolume + (float64(t.Volume) * float64(contents.Amount))
		}

		for srcPinID, _ := range srcMap {
			srcPin := c.pinMap[srcPinID]
			if intArrayContains(extractorTypes, srcPin.TypeId) {
				// if the extractor is not running, skip it
				if pin.ExpiryTime.Before(time.Now()) {
					continue
				}

				t, err := c.clients.DB.GetSDEType(pin.ExtractorDetails.ProductTypeId)
				if err != nil {
					return zeroTime, err
				}

				// find the start time and end time
				//
				// pin.InstallTime
				// find the amount of time left
				totalExtractionTime := pin.ExpiryTime.Sub(pin.InstallTime)
				middleTime := pin.InstallTime.Add(time.Duration(totalExtractionTime.Nanoseconds() / 2))
				durationUntilMiddle := middleTime.Sub(time.Now())
				qtyPerCycle := actualQtyPerCycle(pin.ExtractorDetails)
				cycleDuration := time.Duration(pin.ExtractorDetails.CycleTime) * time.Second
				cyclesUntilMiddle := int32(durationUntilMiddle / cycleDuration)
				remainingExtraction := int32(float32(cyclesUntilMiddle) * float32(qtyPerCycle) * float32(1.5))
				remainingExtractionVolume := float64(remainingExtraction) * float64(t.Volume)

				if currentContentVolume+remainingExtractionVolume > capacity {
					// This is hack because we don't actually know when it'll fill up, just
					// that it will happen before the current extraction finishes
					return time.Now(), nil
				}
			}

			if intArrayContains(bifTypes, srcPin.TypeId) {
				fillRate = fillRate + 7.6
			}
		}

		remainingCapacity := capacity - currentContentVolume
		periodsRemaining := remainingCapacity / fillRate
		minutesRemaining := periodsRemaining * 30
		lpFullAt := c.result.Planet.LastUpdate.Add(time.Duration(minutesRemaining) * time.Minute).Truncate(time.Minute)

		if lpFullAt.Before(earliestFullAt) {
			earliestFullAt = lpFullAt
		}
	}

	return earliestFullAt, nil
}

func actualQtyPerCycle(d esi.GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails) int32 {
	// Found using polinomial regression, has +/- 2% error. Likely overfit
	r := float64(d.HeadRadius)
	ratio := 262306822.5*math.Pow(r, 4) - 20816102.48*math.Pow(r, 3) + 605741.0787*math.Pow(r, 2) - 7441.81751*r + 33.38153241
	return int32(ratio * float64(d.QtyPerCycle))
}

func (c *planetCalculator) nextAttention() (time.Time, error) {
	var zeroTime time.Time
	nextAttention := time.Now().Add(time.Hour * time.Duration(10000)).Truncate(time.Minute)

	if len(c.result.Details.Pins) == 1 {
		return zeroTime, nil
	}

	for _, pin := range c.result.Details.Pins {
		if intArrayContains(extractorTypes, pin.TypeId) {
			if pin.ExpiryTime.Before(nextAttention) {
				nextAttention = pin.ExpiryTime
			}

			continue
		}

		// Deadline calculation for AIFs
		// Logic shortcuts:
		// * AIFs
		// * Only one schematic per launchpad
		// * All the same amount of stuff per cycle for each input
		if intArrayContains(launchpadTypes, pin.TypeId) {
			dstMap := c.sourceRouteMap[pin.PinId]
			ratePerHour := int64(0)
			fewestContents := int64(999999999999)
			fewestContentsId := int32(0)
			contentMap := make(map[int32]int64)
			schematicID := int32(0)

			for _, contents := range pin.Contents {
				contentMap[contents.TypeId] = contents.Amount
			}

			for dstPinID, route := range dstMap {
				dstPin := c.pinMap[dstPinID]
				if intArrayContains(aifTypes, dstPin.TypeId) && contentMap[route.ContentTypeId] < fewestContents {
					fewestContents = contentMap[route.ContentTypeId]
					fewestContentsId = route.ContentTypeId
					schematicID = dstPin.SchematicId
				}
			}

			if schematicID == 0 {
				continue
			}

			schematicQuantity, err := c.clients.DB.GetPlanetarySchematicInputQuantity(schematicID, fewestContentsId)
			if err != nil {
				return zeroTime, err
			}

			var lastCycle time.Time
			for dstPinID, route := range dstMap {
				dstPin := c.pinMap[dstPinID]
				if route.ContentTypeId == fewestContentsId {
					ratePerHour = ratePerHour + schematicQuantity
					if lastCycle.Before(dstPin.LastCycleStart) {
						lastCycle = dstPin.LastCycleStart
					}
				}
			}

			if ratePerHour > 0 {
				hoursRemaining := int64(fewestContents / ratePerHour)
				lastCycle = lastCycle.Add(time.Hour * time.Duration(hoursRemaining+1))
				if lastCycle.Before(nextAttention) {
					nextAttention = lastCycle
				}
			}
		}
	}

	return nextAttention, nil
}

func (c *planetCalculator) QtyPerHour() int32 {
	var qtyPerHour int32 = 0

	for _, pin := range c.result.Details.Pins {
		if intArrayContains(extractorTypes, pin.TypeId) {
			d := pin.ExtractorDetails

			if d.CycleTime == 0 {
				continue
			}

			qtyPerCycle := actualQtyPerCycle(d)
			qtyPerHour = qtyPerHour + (qtyPerCycle * (3600 / d.CycleTime))
		}
	}

	return qtyPerHour
}

func (c *planetCalculator) buildPlanetObservation() (*planetObservation, error) {
	data := c.result.Details
	obs := &planetObservation{
		UpgradeLevel: c.result.Planet.UpgradeLevel,
	}

	for _, pin := range data.Pins {
		if intArrayContains(bifTypes, pin.TypeId) {
			obs.BasicFactories = obs.BasicFactories + 1
			if obs.BasicFactoryOutputID == 0 {
				id, err := c.clients.DB.GetPlanetarySchematicOutputID(pin.SchematicId)
				if err != nil {
					return nil, err
				}

				obs.BasicFactoryOutputID = id
			}
		}

		if intArrayContains(extractorTypes, pin.TypeId) {
			obs.Extractors = obs.Extractors + 1
			obs.QtyPerCycle = obs.QtyPerCycle + pin.ExtractorDetails.QtyPerCycle
			obs.ExtractorHeads = obs.ExtractorHeads + int32(len(pin.ExtractorDetails.Heads))

			if obs.CycleTime != pin.ExtractorDetails.CycleTime {
				if obs.CycleTime != 0 {
					return nil, errors.Errorf("inconsistent extractor cycle times, %d != %d", obs.CycleTime, pin.ExtractorDetails.CycleTime)
				}

				obs.CycleTime = pin.ExtractorDetails.CycleTime
			}

			if obs.ExtractionTypeID != pin.ExtractorDetails.ProductTypeId {
				if obs.ExtractionTypeID != 0 {
					return nil, errors.Errorf("inconsistent extraction types, %d != %d", obs.ExtractionTypeID, pin.ExtractorDetails.ProductTypeId)
				}

				obs.ExtractionTypeID = pin.ExtractorDetails.ProductTypeId
			}

			if obs.HeadRadius != pin.ExtractorDetails.HeadRadius {
				if obs.HeadRadius != 0 {
					if math.Abs(float64(obs.HeadRadius-pin.ExtractorDetails.HeadRadius)) > 0.00006 {
						return nil, errors.Errorf("inconsistent head radius, %f != %f", obs.HeadRadius, pin.ExtractorDetails.HeadRadius)
					}
				} else {
					obs.HeadRadius = pin.ExtractorDetails.HeadRadius
				}
			}
		}
	}

	return obs, nil
}
