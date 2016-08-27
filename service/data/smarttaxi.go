package data

type SmartTaxiDriver struct {
	DriverID    int `json:"DriverId"`
	Lat         float64
	Lng         float64
	CompanyName string
}

type SmartResponse struct {
	Data []SmartTaxiDriver
}
