package car

import "time"

type Record struct {
	ID       string `json:"id"`
	Make     string `json:"make"`
	Model    string `json:"model"`
	Category string `json:"category"`
	Package  string `json:"package"`
	Color    string `json:"color"`
	Year     int    `json:"year"`
	Mileage  int    `json:"mileage"`
	Price    int    `json:"price"`
}

func (c Record) Validate() error {
	if c.ID == "" {
		return ErrorFieldMissing{"ID"}
	}
	if c.Make == "" {
		return ErrorFieldMissing{"Make"}
	}
	if c.Model == "" {
		return ErrorFieldMissing{"Model"}
	}
	if c.Category == "" {
		return ErrorFieldMissing{"Category"}
	}
	if c.Package == "" {
		return ErrorFieldMissing{"Package"}
	}
	if c.Color == "" {
		return ErrorFieldMissing{"Color"}
	}
	if c.Year < 1900 || c.Year > time.Now().Year() {
		return ErrorFieldInvalid{"Year", c.Year}
	}
	if c.Mileage < 0 {
		return ErrorFieldInvalid{"Mileage", c.Mileage}
	}
	if c.Price <= 0 {
		return ErrorFieldInvalid{"Price", c.Price}
	}
	return nil
}
