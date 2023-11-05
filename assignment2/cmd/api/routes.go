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
	router.HandlerFunc(http.MethodGet, "/v1/plantseed", app.listPlantseedHandler)
	router.HandlerFunc(http.MethodPost, "/v1/plantseed", app.createPlantseedHandler)
	router.HandlerFunc(http.MethodGet, "/v1/plantseed/:id", app.showPlantseedHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/plantseed/:id", app.updatePlantseedHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/plantseed/:id", app.deletePlantseedHandler)

	return app.recoverPanic(app.rateLimit(router))
}
