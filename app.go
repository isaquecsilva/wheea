package main

import (
	"context"
	"wheea/backend/core/app"
	"wheea/backend/core/entities"
)

type WheeaApp struct {
	qws *app.WeatherQueryService
	qps *app.QueryPlaceService
	ctx context.Context
}

// NewApp creates a new App application struct
func NewWheeaApp(qps *app.QueryPlaceService, qws *app.WeatherQueryService) *WheeaApp {
	return &WheeaApp{
		qws: qws,
		qps: qps,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *WheeaApp) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *WheeaApp) QueryPlaceBind(cityname string) entities.QueryPlaceApiResponse {
	queryplace := entities.NewQueryPlace(cityname)
	return a.qps.Execute(queryplace)
}

func (a *WheeaApp) WeatherQueryBind(lat, lon float32) entities.WheatherQueryApiResponse {
	wheatherquery := entities.NewWheatherQuery(lat, lon)
	return a.qws.Query(wheatherquery)
}
