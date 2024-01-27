package main

import (
	"context"
)

type WheeaApp struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewWheeaApp() *WheeaApp {
	return &WheeaApp{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *WheeaApp) startup(ctx context.Context) {
	a.ctx = ctx
}