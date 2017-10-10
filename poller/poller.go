package poller

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/curzonj/eve-dwh-golang/types"
)

type Cfg struct {
	RegionID     int32         `env:"REGION_ID,default=10000002"`
	MarketGroups []int         `env:"MARKET_GROUPS,required"`
	Interval     time.Duration `env:"POLLER_INTERVAL,default=5m"`
}

type poller struct {
	clients types.Clients
	logger  log.FieldLogger
	cfg     Cfg
}

func (p *poller) leadingEdgeTick(d time.Duration, f func() error) {
	p.logger.WithField("at", "start").Info()

	err := f()
	if err != nil {
		p.logger.Error(err)
	}

	for range time.Tick(d) {
		err := f()
		if err != nil {
			p.logger.Error(err)
		}
	}
}
