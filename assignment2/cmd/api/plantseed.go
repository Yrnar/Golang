package main

import (
	"errors"
	"fmt"
	"net/http"

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
	err = app.models.Plantseed.Insert(plantseed)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/plantseed/%d", plantseed.ID))
	err = app.writeJSON(w, http.StatusCreated, envelope{"plantseed": plantseed}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showPlantseedHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	plantseed, err := app.models.Plantseed.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"plantseed": plantseed}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
func (app *application) updatePlantseedHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	plantseed, err := app.models.Plantseed.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	var input struct {
		Name   string `json:"name"`
		Family string `json:"family"`
		Amount int32  `json:"amount"`
		Price  int32  `json:"price"`
	}
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	plantseed.Name = input.Name
	plantseed.Family = input.Family
	plantseed.Amount = input.Amount
	plantseed.Price = input.Price
	v := validator.New()
	if data.ValidateMovie(v, plantseed); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	err = app.models.Plantseed.Update(plantseed)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"plantseed": plantseed}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) deletePlantseedHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	err = app.models.Plantseed.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "plantseed successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
