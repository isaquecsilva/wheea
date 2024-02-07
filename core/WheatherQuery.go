package core

type WheatherQuery struct {
	Lat, Lon float64
}

func NewWheatherQuery(Lat, Lon float64) WheatherQuery {
	return WheatherQuery{
		Lat,
		Lon,
	}
}