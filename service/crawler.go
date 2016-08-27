package service

import (
	"fmt"
	"sync"

	"github.com/gen1us2k/log"
	"github.com/maddevsio/openfreecab-crawler/conf"
)

// Crawler is main struct of daemon
// it stores all services that used by
type Crawler struct {
	config *conf.CrawlerConfig

	services  map[string]Service
	waitGroup sync.WaitGroup

	logger log.Logger
}

// NewCrawler creates and returns new CrawlerInstance
func NewCrawler(config *conf.CrawlerConfig) *Crawler {
	os := new(Crawler)
	os.config = config
	os.logger = log.NewLogger("crawler")
	os.services = make(map[string]Service)
	os.AddService(&NambaService{})
	os.AddService(&SmstaxiService{})
	os.AddService(&PelikanService{})
	os.AddService(&SmartTaxiService{})
	return os
}

// Start starts all services in separate goroutine
func (os *Crawler) Start() error {
	os.logger.Info("Starting storage")
	for _, service := range os.services {
		os.logger.Infof("Initializing: %s\n", service.Name())
		if err := service.Init(os); err != nil {
			return fmt.Errorf("initialization of %q finished with error: %v", service.Name(), err)
		}
		os.waitGroup.Add(1)

		go func(srv Service) {
			defer os.waitGroup.Done()
			os.logger.Infof("running %q service\n", srv.Name())
			if err := srv.Run(); err != nil {
				os.logger.Errorf("error on run %q service, %v", srv.Name(), err)
			}
		}(service)
	}
	return nil
}

// AddService adds service into Crawler.services map
func (os *Crawler) AddService(srv Service) {
	os.services[srv.Name()] = srv

}

// Config returns current instance of StorageConfig
func (os *Crawler) Config() conf.CrawlerConfig {
	return *os.config
}

// Stop stops all services running
func (os *Crawler) Stop() {
	os.logger.Info("Worker is stopping...")
	for _, service := range os.services {
		service.Stop()
	}
}

// WaitStop blocks main thread and waits when all goroutines will be stopped
func (os *Crawler) WaitStop() {
	os.waitGroup.Wait()
}
