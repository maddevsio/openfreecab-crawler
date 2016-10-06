package service

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/gen1us2k/log"
	"github.com/maddevsio/openfreecab-crawler/common"
	"github.com/maddevsio/openfreecab-crawler/service/data"
)

type NambaService struct {
	BaseService

	logger      log.Logger
	c           *Crawler
	companyName string
}

func (n *NambaService) Name() string {
	return "namba_crawler"
}

func (n *NambaService) Init(c *Crawler) error {
	n.c = c
	n.logger = log.NewLogger(n.Name())
	n.companyName = "NambaTaxi"
	return nil
}

func (n *NambaService) Run() error {
	n.updateDrivers()
	for range time.Tick(time.Duration(int64(n.c.Config().UpdateInterval)) * time.Second) {
		n.logger.Info("Requesting data")
		n.updateDrivers()
	}
	return nil
}

func (n *NambaService) updateDrivers() {

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
	err = common.CleanStorage(n.c.Config().StorageRootURL, n.companyName)
	if err != nil {
		n.logger.Errorf("Error while cleaning storage, %v", err)
	}
	for _, driver := range drivers.Drivers {
		if driver.Lat == "0.0" || driver.Lon == "0.0" {
			continue
		}
		lat, err := strconv.ParseFloat(driver.Lat, 64)
		if err != nil {
			n.logger.Errorf("Failed to convert lat to float, %v", err)
		}
		lng, err := strconv.ParseFloat(driver.Lon, 64)
		if err != nil {
			n.logger.Errorf("Failed to convert lon to float, %v", err)
		}
		sd := data.StorageDriver{
			Company: n.companyName,
			Lat:     lat,
			Lon:     lng,
		}
		err = common.SaveDriver(n.c.Config().StorageRootURL, sd)

		if err != nil {
			n.logger.Errorf("Failed to save driver, %v", err)
		}
	}
}
