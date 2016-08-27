package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gen1us2k/log"
	"github.com/maddevsio/openfreecab-crawler/common"
	"github.com/maddevsio/openfreecab-crawler/service/data"
)

type NambaService struct {
	BaseService

	logger log.Logger
	c      *Crawler
}

func (n *NambaService) Name() string {
	return "namba_crawler"
}

func (n *NambaService) Init(c *Crawler) error {
	n.c = c
	n.logger = log.NewLogger(n.Name())
	return nil
}

func (n *NambaService) Run() error {
	for range time.Tick(time.Duration(int64(n.c.Config().UpdateInterval)) * time.Second) {
		n.logger.Info("Requesting data")
		driverData, err := common.MakeRequestAndGetBytes(
			"https://nambataxi.kg/core/drivers/free/",
			"GET",
			nil,
		)
		if err != nil {
			n.logger.Errorf("Got error while requesting data, %v", err)
		}
		var drivers data.NambaResponse
		err = json.Unmarshal(driverData, &drivers)
		if err != nil {
			n.logger.Errorf("Got error while parsing data, %v", err)
		}
		fmt.Println(drivers)

	}
	return nil
}
