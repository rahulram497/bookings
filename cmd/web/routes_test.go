package main

import (
	"testing"

	"github.com/go-chi/chi"
	"github.com/rahulram497/bookings/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing

	default:
		t.Errorf("type is not *chi.Mux, but is %T", v)
	}
}
