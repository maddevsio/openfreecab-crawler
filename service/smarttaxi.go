package service

import (
	"encoding/json"
	"time"

	"github.com/gen1us2k/log"
	"github.com/maddevsio/openfreecab-crawler/common"
	"github.com/maddevsio/openfreecab-crawler/service/data"
	"github.com/maddevsio/openfreecab-crawler/storage"
)

type SmartTaxiService struct {
	BaseService

	logger      log.Logger
	c           *Crawler
	companyName string
	cs          *storage.CompanyStorage
}

func (n *SmartTaxiService) Name() string {
	return "smarttaxi_crawler"
}

func (n *SmartTaxiService) Init(c *Crawler) error {
	n.c = c
	n.logger = log.NewLogger(n.Name())
	n.companyName = "SmartTaxi"
	n.cs = storage.NewCompanyStorage()
	return nil
}

func (n *SmartTaxiService) Run() error {
	n.updateDrivers()
	for range time.Tick(time.Duration(int64(n.c.Config().UpdateInterval)) * time.Second) {
		n.logger.Info("Requesting data")
		n.updateDrivers()
	}
	return nil
}

func (n *SmartTaxiService) updateDrivers() {

	driverData, err := common.MakeRequestAndGetBytes(
		"http://smart-taxi.kg/Home/GetDriversOnline",
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
	n.cs.Lock()
	for _, companyName := range n.cs.Data {
		err := common.CleanStorage(n.c.Config().StorageRootURL, companyName)
		if err != nil {
			n.logger.Errorf("Error while cleaning storage, %v", err)
		}
	}
	n.cs.Unlock()
	for _, driver := range drivers.Data {
		if driver.Lat == 0.0 || driver.Lng == 0.0 || !driver.IsFree {
			continue
		}
		sd := data.StorageDriver{
			Company: driver.CompanyName,
			Lat:     driver.Lat,
			Lon:     driver.Lng,
		}
		n.cs.AddCompany(driver.CompanyName)

		err = common.SaveDriver(n.c.Config().StorageRootURL, sd)

		if err != nil {
			n.logger.Errorf("Failed to save driver, %v", err)
		}
	}
}
