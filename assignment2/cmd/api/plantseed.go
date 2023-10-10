package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.assignment2.com/internal/data"
	"golang.assignment2.com/internal/validator"
)

func (app *application) createPlantseedHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name   string `json:"name"`
		Family string `json:"family"`
		Amount int32  `json:"amount,omitempty"`
		Price  int32  `json:"price,omitempty"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	plantseed := &data.Plantseed{
		Name:   input.Name,
		Family: input.Family,
		Amount: input.Amount,
		Price:  input.Price,
	}
	v := validator.New()
	if data.ValidateMovie(v, plantseed); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showPlantseedHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	plantseed := data.Plantseed{
		ID:        id,
		CreatedAt: time.Now(),
		Name:      "AGERATUM",
		Family:    "Asteraceae",
		Amount:    400,
		Price:     3,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"plantseed": plantseed}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
