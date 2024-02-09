package main

import (
	"embed"

	"wheea/adapters"
	"wheea/core/app"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Instanciating services and adapters
	var queryPlaceAdapter *adapters.QueryPlaceAdapter = adapters.NewQueryPlaceAdapter()
	queryPlaceService := app.NewQueryPlaceService(queryPlaceAdapter)

	// Create an instance of the app structure
	app := NewWheeaApp(&queryPlaceService)

	// Create application with options
	err := wails.Run(&options.App{
		Title:      "wheea",
		MinWidth:   344,
		MinHeight:  370,
		Width:      344,
		Height:     370,
		MaxWidth:   600,
		MaxHeight:  600,
		Fullscreen: false,
		Windows: &windows.Options{
			IsZoomControlEnabled: false,
		},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
