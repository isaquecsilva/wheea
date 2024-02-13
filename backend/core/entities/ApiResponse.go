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
	Result struct {
		Latitude  float32
		Longitude float32
		Today     struct {
			WeatherType    uint
			WindSpeed      uint
			Temperature    int
			MaxTemperature int
			MinTemperature int
			Precipitation  uint
		}
		NextDays [6]struct {
			WeatherType    uint
			MaxTemperature int
			MinTemperature int
			Precipitation  uint
		}
	}
	Error string
}
