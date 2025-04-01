package structs

import "strconv"

type ChamplainShuttle struct {
	DateTime  string `json:"Date_Time"`
	UnitID    string `json:"UnitID"`
	Lat       string `json:"Lat"`
	Lon       string `json:"Lon"`
	Knots     string `json:"Knots"`
	Direction string `json:"Direction"`
}

func (champShuttle ChamplainShuttle) ConvertShuttle(oldShuttle *ChamplainShuttle) *Shuttle {
	lat, _ := strconv.ParseFloat(oldShuttle.Lat, 32)
	lon, _ := strconv.ParseFloat(oldShuttle.Lon, 32)
	knots, _ := strconv.ParseFloat(oldShuttle.Knots, 8)
	direction, _ := strconv.ParseUint(oldShuttle.Direction, 10, 16)
	id, _ := strconv.ParseUint(oldShuttle.UnitID, 10, 32)

	newShuttle := Shuttle{
		ID:        uint32(id),
		Lat:       float32(lat),
		Lon:       float32(lon),
		MPH:       uint8(float32(knots) * 1.15),
		Direction: uint16(direction),
	}
	return &newShuttle
}

type Shuttle struct {
	ID        uint32  `json:"id"`
	Lat       float32 `json:"lat"`
	Lon       float32 `json:"lon"`
	MPH       uint8   `json:"mph"` // a shuttle is never going above 255 mph
	Direction uint16  `json:"direction"`
}
