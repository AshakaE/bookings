package main

import (
	"testing"

	"github.com/ashakae/bookings/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T){
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type){
	case *chi.Mux:
		// do nill
	default:
		t.Errorf("unexpected type %T, not a valid *chi.Mux type", v)
	}
}