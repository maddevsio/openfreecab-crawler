package data

type SmartTaxiDriver struct {
	DriverID    int  `json:"DriverId"`
	IsFree      bool `json:"IsFree"`
	Lat         float64
	Lng         float64
	CompanyName string
}

type SmartResponse struct {
	Data []SmartTaxiDriver
}
