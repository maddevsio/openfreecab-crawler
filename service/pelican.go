package service

import (
	"encoding/json"
	"time"

	"github.com/gen1us2k/log"
	"github.com/maddevsio/openfreecab-crawler/common"
	"github.com/maddevsio/openfreecab-crawler/service/data"
)

type PelikanService struct {
	BaseService

	logger      log.Logger
	c           *Crawler
	companyName string
}

func (n *PelikanService) Name() string {
	return "pelikan_service"
}

func (n *PelikanService) Init(c *Crawler) error {
	n.c = c
	n.logger = log.NewLogger(n.Name())
	n.companyName = "Pelikan"
	return nil
}

func (n *PelikanService) Run() error {
	n.updateDrivers()
	for range time.Tick(time.Duration(int64(n.c.Config().UpdateInterval)) * time.Second) {
		n.logger.Info("Requesting data")
		n.updateDrivers()
	}
	return nil
}

func (n *PelikanService) updateDrivers() {

	driverData, err := common.MakeRequestAndGetBytes(
		"http://pelican.kg/ajax/drivers.php",
		"GET",
		nil,
	)
	if err != nil {
		n.logger.Errorf("Got error while requesting data, %v", err)
	}
	var drivers data.PelicanResponse
	err = json.Unmarshal(driverData, &drivers)
	if err != nil {
		n.logger.Errorf("Got error while parsing data, %v", err)
	}
	err = common.CleanStorage(n.c.Config().StorageRootURL, n.companyName)
	if err != nil {
		n.logger.Errorf("Error while cleaning storage, %v", err)
	}
	for key, item := range drivers.Data {
		if key == "drivers" {
			continue
		}
		driver := item.(map[string]interface{})

		lat := driver["lat"].(float64)
		lng := driver["lng"].(float64)
		available := driver["available"].(float64)
		if lat == 0.00 || lng == 0.0 || available == 1.0 {
			continue
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
