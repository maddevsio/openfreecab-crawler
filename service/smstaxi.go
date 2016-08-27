package service

import (
	"encoding/json"
	"time"

	"github.com/gen1us2k/log"
	"github.com/maddevsio/openfreecab-crawler/common"
	"github.com/maddevsio/openfreecab-crawler/service/data"
)

type SmstaxiService struct {
	BaseService

	logger      log.Logger
	c           *Crawler
	companyName string
}

func (n *SmstaxiService) Name() string {
	return "smstaxi_crawler"
}

func (n *SmstaxiService) Init(c *Crawler) error {
	n.c = c
	n.logger = log.NewLogger(n.Name())
	n.companyName = "SmsTaxi"
	return nil
}

func (n *SmstaxiService) Run() error {
	n.updateDrivers()
	for range time.Tick(time.Duration(int64(n.c.Config().UpdateInterval)) * time.Second) {
		n.logger.Info("Requesting data")
		err := common.CleanStorage(n.c.Config().StorageRootURL, n.companyName)
		if err != nil {
			n.logger.Errorf("Error while cleaning storage, %v", err)
		}
		n.updateDrivers()
	}
	return nil
}

func (n *SmstaxiService) updateDrivers() {

	driverData, err := common.MakeRequestAndGetBytes(
		"http://smstaxi.kg/cars?_=timestamp",
		"GET",
		nil,
	)
	if err != nil {
		n.logger.Errorf("Got error while requesting data, %v", err)
	}
	var drivers []data.SmsDriver
	err = json.Unmarshal(driverData, &drivers)
	if err != nil {
		n.logger.Errorf("Got error while parsing data, %v", err)
	}
	for _, driver := range drivers {
		if driver.Lat == 0.0 || driver.Lng == 0.0 {
			continue
		}
		sd := data.StorageDriver{
			Company: n.companyName,
			Lat:     driver.Lat,
			Lon:     driver.Lng,
		}
		err = common.SaveDriver(n.c.Config().StorageRootURL, sd)

		if err != nil {
			n.logger.Errorf("Failed to save driver, %v", err)
		}
	}
}
