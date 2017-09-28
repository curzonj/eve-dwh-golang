package poller

import "time"

type Cfg struct {
	RegionID     int32         `env:"REGION_ID,default=10000002"`
	MarketGroups []int         `env:"MARKET_GROUPS,required"`
	Interval     time.Duration `env:"POLLER_INTERVAL,default=5m"`
}
