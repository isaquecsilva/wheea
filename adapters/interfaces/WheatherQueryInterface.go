package interfaces

import "wheea/core/entities"

type WheatherQueryInterface interface {
	Query(entities.WheatherQuery) entities.WheatherQueryApiResponse
}