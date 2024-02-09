package entities

type WheatherQuery struct {
	Lat, Lon float32
}

func NewWheatherQuery(Lat, Lon float32) WheatherQuery {
	return WheatherQuery{
		Lat,
		Lon,
	}
}
