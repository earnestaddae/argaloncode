package main

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestRouteExist(t *testing.T) {
	testRoutes := testApp.routes()
	chiRoutes := testRoutes.(chi.Router)

	routeExists(t, chiRoutes, "/v1/healthcheck")
	routeExists(t, chiRoutes, "/v1/divisions")
}

func routeExists(t *testing.T, routes chi.Router, route string) {
	found := false

	_ = chi.Walk(routes, func(_, foundRoute string, _ http.Handler, _ ...func(http.Handler) http.Handler) error {
		if route == foundRoute {
			found = true
		}
		return nil
	})

	if !found {
		t.Errorf("did not find %s in registered routes", route)
	}
}
