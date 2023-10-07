package data

import (
	"time"

	"golang.assignment2.com/internal/validator"
)

type Plantseed struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Name      string    `json:"name"`
	Family    string    `json:"family"`
	Amount    int32     `json:"amount,omitempty"`
	Price     int32     `json:"price,omitempty"`
}

func ValidateMovie(v *validator.Validator, plantseed *Plantseed) {
	v.Check(plantseed.Name != "", "name", "must be provided")
	v.Check(len(plantseed.Name) <= 500, "name", "must not be more than 500 bytes long")
	v.Check(plantseed.Family != "", "family", "must be provided")
	v.Check(len(plantseed.Family) <= 500, "family", "must not be more than 500 bytes long")
	v.Check(plantseed.Amount != 0, "amount", "must be provided")
	v.Check(plantseed.Amount >= 0, "amount", "must be greater than 0")
	v.Check(plantseed.Price != 0, "price", "must be provided")
	v.Check(plantseed.Price >= 0, "price", "must be greater than 0")
}
