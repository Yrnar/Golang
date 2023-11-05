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
		Name   *string `json:"name"`
		Family *string `json:"family"`
		Amount *int32  `json:"amount"`
		Price  *int32  `json:"price"`
	}
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	if input.Name != nil {
		plantseed.Name = *input.Name
	}
	if input.Family != nil {
		plantseed.Family = *input.Family
	}
	if input.Amount != nil {
		plantseed.Amount = *input.Amount
	}
	if input.Price != nil {
		plantseed.Price = *input.Price
	}
	v := validator.New()
	if data.ValidateMovie(v, plantseed); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	err = app.models.Plantseed.Update(plantseed)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
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

func (app *application) listPlantseedHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name   string
		Family string
		Amount int
		Price  int
		data.Filters
	}
	v := validator.New()
	qs := r.URL.Query()
	input.Name = app.readString(qs, "name", "")
	input.Family = app.readString(qs, "family", "")
	input.Amount = app.readInt(qs, "amount", 1, v)
	input.Price = app.readInt(qs, "price", 1, v)

	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "name", "family", "amount", "price", "-id", "-name", "-family", "-amount", "-price"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	plantseeds, metadata, err := app.models.Plantseed.GetAll(input.Name, input.Family, input.Amount, input.Price, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"plantseeds": plantseeds,  "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
