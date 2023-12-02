package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/plantseed", app.requirePermission("plantseed:read", app.listPlantseedHandler))
	router.HandlerFunc(http.MethodPost, "/v1/plantseed", app.requirePermission("plantseed:write", app.createPlantseedHandler))
	router.HandlerFunc(http.MethodGet, "/v1/plantseed/:id", app.requirePermission("plantseed:read", app.showPlantseedHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/plantseed/:id", app.requirePermission("plantseed:write", app.updatePlantseedHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/plantseed/:id", app.requirePermission("plantseed:write", app.deletePlantseedHandler))

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)

	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	return app.recoverPanic(app.rateLimit(app.authenticate(router)))
}

