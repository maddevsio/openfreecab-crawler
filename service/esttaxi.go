package service

import (
	"encoding/json"
	"time"

	"github.com/gen1us2k/log"
	"github.com/maddevsio/openfreecab-crawler/common"
	"github.com/maddevsio/openfreecab-crawler/service/data"
	"github.com/maddevsio/openfreecab-crawler/storage"
)

type EstService struct {
	BaseService

	logger      log.Logger
	c           *Crawler
	companyName string
	cs          *storage.CompanyStorage
}

func (n *EstService) Name() string {
	return "esttaxi_crawler"
}

func (n *EstService) Init(c *Crawler) error {
	n.c = c
	n.logger = log.NewLogger(n.Name())
	n.companyName = "EstTaxi"
	n.cs = storage.NewCompanyStorage()
	return nil
}

func (n *EstService) Run() error {
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

func (n *EstService) updateDrivers() {

	driverData, err := common.MakeRequestAndGetBytes(
		"http://siteapi.estaxi.org/drivers/taxi.geojson?lang=en-US",
		"GET",
		nil,
	)
	if err != nil {
		n.logger.Errorf("Got error while requesting data, %v", err)
	}

	var drivers data.EstResponse
	err = json.Unmarshal(driverData, &drivers)
	if err != nil {
		n.logger.Errorf("Got error while parsing data, %v", err)
	}
	for _, driver := range drivers.Features {
		if driver.Properties.Status != "свободен" {
			continue
		}
		sd := data.StorageDriver{
			Company: n.companyName,
			Lat:     driver.Geometry.Coordinates[0],
			Lon:     driver.Geometry.Coordinates[1],
		}
		n.cs.AddCompany(driver.Properties.TaxiName)
		err = common.SaveDriver(n.c.Config().StorageRootURL, sd)

		if err != nil {
			n.logger.Errorf("Failed to save driver, %v", err)
		}
	}
}
