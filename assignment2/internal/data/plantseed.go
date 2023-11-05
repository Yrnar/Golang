package data

import (
	"database/sql"
	"errors"
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

type PlantseedModel struct {
	DB *sql.DB
}

func (m PlantseedModel) Insert(plantseed *Plantseed) error {
	query := `
		INSERT INTO plantseed (name, family, amount, price)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at,`
	args := []interface{}{plantseed.Name, plantseed.Family, plantseed.Amount, plantseed.Price}
	return m.DB.QueryRow(query, args...).Scan(&plantseed.ID, &plantseed.CreatedAt)

}
func (m PlantseedModel) Get(id int64) (*Plantseed, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
		SELECT id, created_at, name, family, amount, price
		FROM plantseed
		WHERE id = $1`
	var plantseed Plantseed
	err := m.DB.QueryRow(query, id).Scan(
		&plantseed.ID,
		&plantseed.CreatedAt,
		&plantseed.Name,
		&plantseed.Family,
		&plantseed.Amount,
		&plantseed.Price,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &plantseed, nil
}
func (m PlantseedModel) Update(plantseed *Plantseed) error {
	query := `
		UPDATE plantseed
		SET name = $1, family = $2, amount = $3, price = $4
		WHERE id = $5`
	args := []interface{}{
		&plantseed.Name,
		&plantseed.Family,
		&plantseed.Amount,
		&plantseed.Price,
		&plantseed.ID,
	}
	return m.DB.QueryRow(query, args...).Scan()
}
func (m PlantseedModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
		DELETE FROM plantseed
		WHERE id = $1`
	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}
