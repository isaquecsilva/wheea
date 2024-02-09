package adapters

import (
	"testing"
	"wheea/backend/core/entities"

	"github.com/stretchr/testify/assert"
)

func Test_Query(test *testing.T) {
	test.Run("should get a valid weather query response", func(t *testing.T) {
		washington := struct {
			Lat float32
			Lon float32
		}{38.89511, -77.03637}

		weatherquery := entities.NewWheatherQuery(washington.Lat, washington.Lon)
		weatherQueryAdapter := NewWheatherQueryAdapter()
		response := weatherQueryAdapter.Query(weatherquery)

		assert.Emptyf(t, response.Error, "the error should be \"\"(empty), but actual is not: %s.",
			response.Error)
	})
}
