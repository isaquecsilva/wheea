package interfaces

import "wheea/core/entities"

type QueryPlaceInterface interface {
	Query(entities.QueryPlace) entities.QueryPlaceApiResponse
}