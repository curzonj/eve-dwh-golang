package data

type PollerCfg struct {
	RegionID     int32 `env:"REGION_ID,default=10000002"`
	MarketGroups []int `env:"MARKET_GROUPS,required"`
}
