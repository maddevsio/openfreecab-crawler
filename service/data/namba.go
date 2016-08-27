package data

type NambaDriver struct {
	Lat string `json:"lat"`
	Lon string `json:"lng"`
}

type NambaResponse struct {
	Drivers []NambaDriver `json:"drivers"`
}
