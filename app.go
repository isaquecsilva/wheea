package main

import (
	"context"
	"wheea/backend/core/app"
	"wheea/backend/core/entities"
)

type WheeaApp struct {
	qws *app.WheatherQueryService
	qps *app.QueryPlaceService
	ctx context.Context
}

// NewApp creates a new App application struct
func NewWheeaApp(qps *app.QueryPlaceService) *WheeaApp {
	return &WheeaApp{
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

func (a *WheeaApp) QueryWheatherBind(lat, lon float32) entities.WheatherQueryApiResponse {
	wheatherquery := entities.NewWheatherQuery(lat, lon)
	return a.qws.Query(wheatherquery)
}
