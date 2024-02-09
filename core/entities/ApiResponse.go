package entities

type QueryPlaceApiResponse struct {
	Code    int
	Results []struct {
		Name      string
		Latitude  float32
		Longitude float32
		Country   string
	}
	Error string
}

type WheatherQueryApiResponse struct {
	Code   int
	Result []struct {
		Latitude  float32
		Longitude float32
		Today     struct {
			WindSpeed      uint
			MaxTemperature uint
			MinTemperature uint
			Precipitation  uint
		}
		NextDays [6]struct {
			WindSpeed      uint
			MaxTemperature uint
			MinTemperature uint
			Precipitation  uint
		}
	}
	Error string
}
