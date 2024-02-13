package app

import (
	"wheea/backend/adapters/interfaces"
	"wheea/backend/core/entities"
)

type QueryPlaceService struct {
	qpi interfaces.QueryPlaceInterface
}

func NewQueryPlaceService(qpi interfaces.QueryPlaceInterface) QueryPlaceService {
	return QueryPlaceService{
		qpi,
	}
}

func (qps QueryPlaceService) Execute(qp entities.QueryPlace) entities.QueryPlaceApiResponse {
	return qps.qpi.Query(qp)
}
