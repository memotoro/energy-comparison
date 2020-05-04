package data

import (
	"fmt"
	"math"
)

// EnergyPlan -
type EnergyPlan struct {
	Supplier       string     `json:"supplier"`
	Plan           string     `json:"plan"`
	Rates          []Rate     `json:"rates"`
	StandingCharge *float64   `json:"standing_charge"`
	Discount       []Discount `json:"discounts"`
}

// Discount -
type Discount struct {
	AppliesTo string  `json:"applies_to"`
	Cap       *int    `json:"cap"`
	Value     float64 `json:"value"`
	ValueType string  `json:"value_type"`
}

// Rate -
type Rate struct {
	Price     float64 `json:"price"`
	Threshold *int    `json:"threshold"`
}

// CustomerPlan -
type CustomerPlan struct {
	EnergyPlan EnergyPlan
	Total      float64
}

// TotalDisplay -
func (cp *CustomerPlan) TotalDisplay() string {
	return fmt.Sprintf("%.2f", math.Round(cp.Total)/100)
}
