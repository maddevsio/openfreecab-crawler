package data

type EstDriver struct {
	Type     string `json:"type"`
	ID       int    `json:"id"`
	Geometry struct {
		Type        string `json:"type"`
		Coordinates []float64
	} `json:"geometry"`
	Properties struct {
		TaxiName  string `json:"taxi_name"`
		CarName   string `json:"car_name"`
		Timestamp int    `json:"timestamp"`

		Status   string `json:"status"`
		StatusID int    `json:"status_id"`
	} `json:"properties"`
}

type EstResponse struct {
	Features []EstDriver `json:"features"`
}
