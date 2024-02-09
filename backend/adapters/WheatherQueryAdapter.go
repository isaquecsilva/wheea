package adapters

import (
	"encoding/json"
	"net/http"
	"wheea/backend/core/entities"

	"github.com/innotechdevops/openmeteo"
)

type WheatherQueryAdapter struct {
	meteo  openmeteo.OpenMeteo
	params openmeteo.Parameter
}

func NewWheatherQueryAdapter() *WheatherQueryAdapter {
	return &WheatherQueryAdapter{
		meteo: openmeteo.New(),
		params: openmeteo.Parameter{
			CurrentWeather: openmeteo.Bool(true),
			Daily: &[]string{
				openmeteo.DailyTemperature2mMax,
				openmeteo.DailyTemperature2mMin,
				openmeteo.DailyWeatherCode,
				openmeteo.DailyPrecipitationProbabilityMean,
			},
		},
	}
}

func (wqa *WheatherQueryAdapter) Query(wq entities.WheatherQuery) entities.WheatherQueryApiResponse {
	wqa.params.Latitude = openmeteo.Float32(wq.Lat)
	wqa.params.Longitude = openmeteo.Float32(wq.Lon)

	if rawresponse, err := wqa.meteo.Execute(wqa.params); err != nil {
		return entities.WheatherQueryApiResponse{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		}
	} else {
		response := entities.WheatherQueryApiResponse{
			Code:  http.StatusOK,
			Error: "",
		}

		var apiresmap map[string]any
		switch err := json.Unmarshal([]byte(rawresponse), &apiresmap); err {
		case nil:
			response.Result = struct {
				Latitude  float32
				Longitude float32
				Today     struct {
					WeatherType    uint
					WindSpeed      uint
					Temperature    int
					MaxTemperature int
					MinTemperature int
					Precipitation  uint
				}
				NextDays [6]struct {
					WeatherType    uint
					MaxTemperature int
					MinTemperature int
					Precipitation  uint
				}
			}{
				Latitude:  float32(apiresmap["latitude"].(float64)),
				Longitude: float32(apiresmap["longitude"].(float64)),
				Today: struct {
					WeatherType    uint
					WindSpeed      uint
					Temperature    int
					MaxTemperature int
					MinTemperature int
					Precipitation  uint
				}{
					WeatherType:    uint(apiresmap["current_weather"].(map[string]interface{})["weathercode"].(float64)),
					WindSpeed:      uint(apiresmap["current_weather"].(map[string]interface{})["windspeed"].(float64)),
					Temperature:    int(apiresmap["current_weather"].(map[string]interface{})["temperature"].(float64)),
					MaxTemperature: int(apiresmap["daily"].(map[string]any)["temperature_2m_max"].([]any)[0].(float64)),
					MinTemperature: int(apiresmap["daily"].(map[string]any)["temperature_2m_min"].([]any)[0].(float64)),
					Precipitation:  uint(apiresmap["daily"].(map[string]any)["precipitation_probability_mean"].([]any)[0].(float64)),
				},
			}

			for index, _ := range response.Result.NextDays {
				response.Result.NextDays[index] = struct {
					WeatherType    uint
					MaxTemperature int
					MinTemperature int
					Precipitation  uint
				}{
					WeatherType:    uint(apiresmap["daily"].(map[string]interface{})["weathercode"].([]any)[index].(float64)),
					MaxTemperature: int(apiresmap["daily"].(map[string]interface{})["temperature_2m_max"].([]any)[index].(float64)),
					MinTemperature: int(apiresmap["daily"].(map[string]interface{})["temperature_2m_min"].([]any)[index].(float64)),
					Precipitation:  uint(apiresmap["daily"].(map[string]interface{})["precipitation_probability_mean"].([]any)[index].(float64)),
				}
			}
		default:
			response.Code = http.StatusInternalServerError
			response.Error = err.Error()
		}

		return response

	}
}
