package adapters

import (
	"net/http"
	"testing"
	"wheea/core/entities"

	"github.com/stretchr/testify/assert"
)

func TestQuery(test *testing.T) {
	test.Run("should get a valid api response", func(t *testing.T) {
		qpadapter := NewQueryPlaceAdapter("https://geocoding-api.open-meteo.com/v1/search?name=%s")
		queryplace := entities.NewQueryPlace("Berlin")
		response := qpadapter.Query(queryplace)

		assert.Equalf(t, http.StatusOK, response.Code, "expected valid api response (success), but status code is: %d. Error: %s", response.Code, response.Error)

		assert.Greater(t, len(response.Results), 0, "expected 'results' key to have a greater length than 0")
	})
}
