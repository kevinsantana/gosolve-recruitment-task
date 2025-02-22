package server

import (
	"net/http"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"

	"github.com/kevinsantana/gosolve-recruitment-task/internal/rest"
	"github.com/kevinsantana/gosolve-recruitment-task/internal/rest/handlers"
	"github.com/kevinsantana/gosolve-recruitment-task/internal/rest/middlewares"
)

type Routes []Route

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc fiber.Handler
	Public      bool
	Scopes      []string
}

func Router(health rest.HealthWebHandler) *fiber.App {
	var healthCheck = Routes{
		{
			Name:        "Healthcheck",
			Method:      http.MethodGet,
			Pattern:     "/healthcheck",
			HandlerFunc: health.Liveness,
			Public:      true,
		},
		{
			Name:        "Readiness",
			Method:      http.MethodGet,
			Pattern:     "/readiness",
			HandlerFunc: health.Readiness,
			Public:      true,
		},
	}

	var searchValue = Routes{
		{
			Name:        "Search index by value",
			Method:      http.MethodGet,
			Pattern:     "/:value/search",
			HandlerFunc: handlers.SearchByValueHandler,
			Public:      true,
		},
	}

	r := fiber.New(fiber.Config{
		Prefork:               false,
		CaseSensitive:         false,
		StrictRouting:         false,
		ServerHeader:          "*",
		AppName:               "Gosolve Recruitment Task",
		Immutable:             true,
		DisableStartupMessage: true,
		ErrorHandler:          middlewares.ErrorHandler(),
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
	})

	api := r.Group("/")
	for _, route := range healthCheck {
		api.Add(route.Method, route.Pattern, route.HandlerFunc)
	}

	v1 := api.Group("/api/v1")

	var routes []Route
	routes = append(routes, searchValue...)

	for _, route := range routes {
		v1.Add(route.Method, route.Pattern, route.HandlerFunc)
	}

	r.Use(middlewares.RouteNotFound())

	return r
}
