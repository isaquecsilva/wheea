package app

import (
	"wheea/core/entities"
	"wheea/adapters/interfaces"
)

type WheatherQueryService struct {
	wqi interfaces.WheatherQueryInterface
}

func NewWheatherQueryService(wqi interfaces.WheatherQueryInterface) WheatherQueryService {
	return WheatherQueryService{
		wqi,
	}
}

func (wqs WheatherQueryService) Query(wq entities.WheatherQuery) entities.WheatherQueryApiResponse {
	return wqs.wqi.Query(wq)
}