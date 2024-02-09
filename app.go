package main

import (
	"context"
	"wheea/core/app"
	"wheea/core/entities"
)

type WheeaApp struct {
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

func (a *WheeaApp) QueryWheather(cityname string) entities.QueryPlaceApiResponse {

	queryplace := entities.NewQueryPlace(cityname)
	return a.qps.Execute(queryplace)
}
