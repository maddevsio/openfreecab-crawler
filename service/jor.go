package service

import (
	"encoding/json"
	"time"

	"github.com/gen1us2k/log"
	"github.com/maddevsio/openfreecab-crawler/common"
	"github.com/maddevsio/openfreecab-crawler/service/data"
	"github.com/maddevsio/openfreecab-crawler/storage"
)

type JorgoTaxiService struct {
	BaseService

	logger      log.Logger
	c           *Crawler
	companyName string
	cs          *storage.CompanyStorage
}

func (n *JorgoTaxiService) Name() string {
	return "jorgo_crawler"
}

func (n *JorgoTaxiService) Init(c *Crawler) error {
	n.c = c
	n.logger = log.NewLogger(n.Name())
	n.companyName = "Jorgo"
	n.cs = storage.NewCompanyStorage()
	return nil
}

func (n *JorgoTaxiService) Run() error {
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

func (n *JorgoTaxiService) updateDrivers() {

	driverData, err := common.MakeRequestAndGetBytes(
		"http://jorgo.smart-taxi.kg/Home/GetDriversOnline",
		"POST",
		nil,
	)
	if err != nil {
		n.logger.Errorf("Got error while requesting data, %v", err)
	}
	var drivers data.SmartResponse
	err = json.Unmarshal(driverData, &drivers)
	if err != nil {
		n.logger.Errorf("Got error while parsing data, %v", err)
	}
	for _, driver := range drivers.Data {
		if driver.Lat == 0.0 || driver.Lng == 0.0 || !driver.IsFree {
			continue
		}
		sd := data.StorageDriver{
			Company: driver.CompanyName,
			Lat:     driver.Lat,
			Lon:     driver.Lng,
		}

		err = common.SaveDriver(n.c.Config().StorageRootURL, sd)

		if err != nil {
			n.logger.Errorf("Failed to save driver, %v", err)
		}
	}
}
