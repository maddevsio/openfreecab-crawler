package conf

// CrawlerConfig stores all configuration of service
type CrawlerConfig struct {
	UpdateInterval int
	StorageRootURL string
	TestMode       bool
}
