package adapters

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"wheea/backend/core/entities"
)

type QueryPlaceAdapter struct {
	endpoint string
}

func NewQueryPlaceAdapter() *QueryPlaceAdapter {
	return &QueryPlaceAdapter{
		"https://geocoding-api.open-meteo.com/v1/search?name=%s",
	}
}

func (qpa *QueryPlaceAdapter) Query(qp entities.QueryPlace) entities.QueryPlaceApiResponse {
	response, err := http.Get(
		fmt.Sprintf(qpa.endpoint, url.QueryEscape(qp.CityName)),
	)

	if err != nil {
		return entities.QueryPlaceApiResponse{
			Code:    http.StatusInternalServerError,
			Results: nil,
			Error:   err.Error(),
		}
	} else {
		defer response.Body.Close()

		var apiresponse map[string]any

		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(&apiresponse); err != nil {
			return entities.QueryPlaceApiResponse{
				Code:  http.StatusInternalServerError,
				Error: err.Error(),
			}
		}

		// Whether we got some error message, the 'reason' key will be set in the json,
		// and we can retrieve its message
		if err, ok := apiresponse["reason"]; ok {
			return entities.QueryPlaceApiResponse{
				Code:  response.StatusCode,
				Error: err.(string),
			}
		} else {
			var finalApiResponse entities.QueryPlaceApiResponse

			// If everything is ok...
			if results, ok := apiresponse["results"]; ok {
				apiresponse = nil

				for _, m := range results.([]interface{}) {
					m := m.(map[string]interface{})

					if country, ok := m["country"]; ok {
						finalApiResponse.Results = append(finalApiResponse.Results, struct {
							Name      string
							Latitude  float32
							Longitude float32
							Country   string
						}{
							m["name"].(string),
							float32(m["latitude"].(float64)),
							float32(m["longitude"].(float64)),
							country.(string),
						})
					}
				}

				finalApiResponse.Code = response.StatusCode
				return finalApiResponse
			} else {
				return entities.QueryPlaceApiResponse{
					Code:    http.StatusBadRequest,
					Results: nil,
					Error:   "couldn't fetch city's informations: invalid city name.",
				}
			}

		}
	}
}
