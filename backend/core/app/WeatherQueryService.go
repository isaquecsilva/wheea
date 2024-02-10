package app

import (
	"wheea/backend/adapters/interfaces"
	"wheea/backend/core/entities"
)

type WeatherQueryService struct {
	wqi interfaces.WheatherQueryInterface
}

func NewWeatherQueryService(wqi interfaces.WheatherQueryInterface) WeatherQueryService {
	return WeatherQueryService{
		wqi,
	}
}

func (wqs WeatherQueryService) Query(wq entities.WheatherQuery) entities.WheatherQueryApiResponse {
	return wqs.wqi.Query(wq)
}
