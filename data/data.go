package data

import (
	"fmt"
	"math"
)

// EnergyPlan -
type EnergyPlan struct {
	Supplier       string `json:"supplier"`
	Plan           string `json:"plan"`
	Rates          []Rate `json:"rates"`
	StandingCharge *int   `json:"standing_charge"`
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
	return fmt.Sprintf("%.2f", math.Ceil(cp.Total)/100)
}
