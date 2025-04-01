package structs

type ChamplainShuttle struct {
	DateTime  string `json:"Date_Time"`
	UnitID    string `json:"UnitID"`
	Lat       string `json:"Lat"`
	Lon       string `json:"Lon"`
	Knots     string `json:"Knots"`
	Direction string `json:"Direction"`
}

type Shuttle struct {
	ID        int     `json:"id"`
	Lat       float32 `json:"lat"`
	Lon       float32 `json:"lon"`
	MPH       int64   `json:"mph"`
	Direction int64   `json:"direction"`
}
