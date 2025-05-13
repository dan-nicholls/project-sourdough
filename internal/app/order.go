package app

import (
	"fmt"
	"regexp"
	"slices"
	"time"
)

type BreadOptions struct {
	Shape string
	Flour string
	Topping string
	Scoring string
}

func (b *BreadOptions) Validate() error {
	if !slices.Contains(availableConfigs.ShapeTypes, b.Shape) {
		return fmt.Errorf("invalid shape: %q is not in %v", b.Shape, availableConfigs.ShapeTypes)
	}
	if !slices.Contains(availableConfigs.FlourTypes, b.Flour) {
		return fmt.Errorf("invalid flour: %q is not in %v", b.Flour, availableConfigs.FlourTypes)
	}
	if !slices.Contains(availableConfigs.ToppingTypes, b.Topping) {
		return fmt.Errorf("invalid topping: %q is not in %v", b.Topping, availableConfigs.ToppingTypes)
	}
	if !slices.Contains(availableConfigs.ScoringTypes, b.Scoring) {
		return fmt.Errorf("invalid scoring: %q is not in %v", b.Scoring, availableConfigs.ScoringTypes)
	}
	return nil
}


type Order struct {
	Id int64
	Name string
	Email string
	OrderDate time.Time
	Status string
	Options BreadOptions
}

func (a *AppService) InsertOrder(o Order) (Order, error) {
	query := `
		INSERT INTO orders (name, email, shape, flour, topping, scoring, order_date, status) 
		VALUES (?,?,?,?,?,?,?,?)
	`

	result, err := a.Db.Exec(query, 
		o.Name, 
		o.Email,
		o.Options.Shape,
		o.Options.Flour,
		o.Options.Topping,
		o.Options.Scoring,
		o.OrderDate,
		o.Status,
	)
	
	if err != nil {
		return Order{}, fmt.Errorf("failed to insert order: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Order{}, fmt.Errorf("failed to get last inserted id: %w", err)
	}

	o.Id = id
	return o, nil
}

func isValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func (o *Order) Validate() error {
	if len(o.Name) <= 0 {
		return fmt.Errorf("Name cannot be none")	
	}
	if !isValidEmail(o.Email) {
		return fmt.Errorf("Email is not valid")
	}

	err := o.Options.Validate()
	if err != nil {
		return err
	}

	return nil
}
