package interfaces

import "wheea/backend/core/entities"

type WheatherQueryInterface interface {
	Query(entities.WheatherQuery) entities.WheatherQueryApiResponse
}
