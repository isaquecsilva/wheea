package adapters

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"wheea/core/entities"
)

type QueryPlaceAdapter struct {
	endpoint string
}

func NewQueryPlaceAdapter(endpoint string) *QueryPlaceAdapter {
	return &QueryPlaceAdapter{
		endpoint,
	}
}

func (qpa *QueryPlaceAdapter) Query(qp entities.QueryPlace) entities.QueryPlaceApiResponse {
	response, err := http.Get(
		fmt.Sprintf(qpa.endpoint, qp.CityName),
	)

	if err != nil {
		return entities.QueryPlaceApiResponse{
			Code:    http.StatusInternalServerError,
			Results: nil,
			Error:   err,
		}
	} else {
		defer response.Body.Close()

		var apiresponse map[string]any

		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(&apiresponse); err != nil {
			return entities.QueryPlaceApiResponse{
				Code:  http.StatusInternalServerError,
				Error: err,
			}
		}

		// Whether we got some error message, the 'reason' key will be set in the json,
		// and we can retrieve its message
		if err, ok := apiresponse["reason"]; ok {
			return entities.QueryPlaceApiResponse{
				Code:  response.StatusCode,
				Error: errors.New(err.(string)),
			}
		} else {
			var finalApiResponse entities.QueryPlaceApiResponse

			// If everything is ok...
			for _, m := range apiresponse["results"].([]interface{}) {
				m := m.(map[string]interface{})
				finalApiResponse.Results = append(finalApiResponse.Results, struct {
					Name      string
					Latitude  float32
					Longitude float32
					Country   string
				}{
					m["name"].(string),
					float32(m["latitude"].(float64)),
					float32(m["longitude"].(float64)),
					m["country"].(string),
				})
			}

			finalApiResponse.Code = response.StatusCode
			finalApiResponse.Error = nil
			return finalApiResponse
		}
	}
}
