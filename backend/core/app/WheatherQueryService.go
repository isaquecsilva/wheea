package app

import (
	"wheea/backend/adapters/interfaces"
	"wheea/backend/core/entities"
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
