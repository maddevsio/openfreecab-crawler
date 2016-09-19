package storage

import "sync"

type CompanyStorage struct {
	sync.RWMutex
	Data map[string]string
}

func NewCompanyStorage() *CompanyStorage {
	c := new(CompanyStorage)
	c.Data = make(map[string]string)
	return c
}

func (c *CompanyStorage) AddCompany(companyName string) {
	c.Lock()
	c.Data[companyName] = companyName
	c.Unlock()
}
