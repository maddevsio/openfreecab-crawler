package conf

import (
	"os"

	"github.com/gen1us2k/log"
	"github.com/urfave/cli"
)

// Version stores current service version
var (
	Version        string
	StorageRootURL string
	UpdateInterval int
	LogLevel       string
	TestMode       bool
)

type Configuration struct {
	data *CrawlerConfig
	app  *cli.App
}

// NewConfigurator is constructor and creates a new copy of Configuration
func NewConfigurator() *Configuration {
	Version = "0.1dev"
	app := cli.NewApp()
	app.Name = "Open free cab crawler"
	app.Usage = "Crawl drivers from different sources"
	return &Configuration{
		data: &CrawlerConfig{},
		app:  app,
	}
}

func (c *Configuration) fillConfig() *CrawlerConfig {
	return &CrawlerConfig{
		StorageRootURL: StorageRootURL,
		UpdateInterval: UpdateInterval,
		TestMode:       TestMode,
	}
}

// Run is wrapper around cli.App
func (c *Configuration) Run() error {
	c.app.Before = func(ctx *cli.Context) error {
		log.SetLevel(log.MustParseLevel(LogLevel))
		return nil
	}
	c.app.Flags = c.setupFlags()
	return c.app.Run(os.Args)
}

// App is public method for Configuration.app
func (c *Configuration) App() *cli.App {
	return c.app
}

func (c *Configuration) setupFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "storage_root_url",
			Value:       "http://localhost:8090",
			Usage:       "OpenfreeCabStorage root url",
			EnvVar:      "OPEN_FREE_CAB_STORAGE_URL",
			Destination: &StorageRootURL,
		},
		cli.StringFlag{
			Name:        "loglevel",
			Value:       "debug",
			Usage:       "set log level",
			Destination: &LogLevel,
			EnvVar:      "LOG_LEVEL",
		},
		cli.BoolFlag{
			Name:        "test_mode",
			Usage:       "set test mode",
			Destination: &TestMode,
			EnvVar:      "TEST_MODE",
		},
		cli.IntFlag{
			Name:        "update_interval",
			Usage:       "Set update interval",
			Value:       15,
			Destination: &UpdateInterval,
			EnvVar:      "UPDATE_INTERVAL",
		},
	}
}

// Get returns filled CrawlerConfig
func (c *Configuration) Get() *CrawlerConfig {
	c.data = c.fillConfig()
	return c.data
}
