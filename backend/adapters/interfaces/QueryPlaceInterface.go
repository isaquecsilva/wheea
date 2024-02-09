package interfaces

import "wheea/backend/core/entities"

type QueryPlaceInterface interface {
	Query(entities.QueryPlace) entities.QueryPlaceApiResponse
}
