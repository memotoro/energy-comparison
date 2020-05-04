package calculations

import (
	"fmt"

	"uswitch.com/comparison/data"
)

// Calculator -
type Calculator interface {
	CalculatePlansForCustomer(plans []data.EnergyPlan, yearlyUsage int) []data.CustomerPlan
	CalculateEnergyUsedAnnually(plans []data.EnergyPlan, supplierName, planName string, monthlySpend float64) (*float64, error)
}

type calculator struct {
	taxRate float64
}

// NewDefaultCalculator -
func NewDefaultCalculator(taxRate float64) Calculator {
	return &calculator{taxRate: taxRate}
}

// CalculatePlansForCustomer -
func (c *calculator) CalculatePlansForCustomer(plans []data.EnergyPlan, yearlyUsage int) []data.CustomerPlan {
	customerPlans := make([]data.CustomerPlan, 0)

	for _, plan := range plans {
		customerUsage := yearlyUsage
		cumulative := float64(0)

		for _, rate := range plan.Rates {
			if rate.Threshold != nil {
				// Checks if the threshold is included in the usage
				if customerUsage >= *rate.Threshold {
					cumulative += float64(*rate.Threshold) * rate.Price
					// Calculates the remaining usage for next rate
					customerUsage -= *rate.Threshold
				} else {
					cumulative += float64(customerUsage) * rate.Price
					customerUsage = 0
				}
			} else {
				value := float64(customerUsage) * rate.Price
				cumulative += value
			}
		}

		if plan.StandingCharge != nil {
			// Adds the standing charge for one year
			cumulative += float64(*plan.StandingCharge * 365)
		}

		// Apply discounts
		total := applyDiscount(cumulative, plan)

		// Adds taxes
		total = addTax(total, c.taxRate)
		customerPlans = append(customerPlans, data.CustomerPlan{EnergyPlan: plan, Total: total})
	}

	return customerPlans
}

// CalculateEnergyUsedAnnually -
func (c *calculator) CalculateEnergyUsedAnnually(plans []data.EnergyPlan, supplierName, planName string, monthlySpend float64) (*float64, error) {
	var selectedPlan *data.EnergyPlan

	for _, plan := range plans {
		// Selects the plan if matches the criteria
		if plan.Supplier == supplierName && plan.Plan == planName {
			selectedPlan = &plan
			break
		}
	}

	if selectedPlan == nil {
		return nil, fmt.Errorf("Not plan available for supplier %s and plan %s", supplierName, planName)
	}

	if monthlySpend == 0 {
		noUsage := float64(0)
		return &noUsage, nil
	}

	yearlySpend := float64(monthlySpend * 100 * 12)
	total := removeTax(yearlySpend, c.taxRate)

	total = applyDiscount(total, *selectedPlan)

	if selectedPlan.StandingCharge != nil {
		// Removes the standing charge for the year
		total -= float64(*selectedPlan.StandingCharge * 365)
	}

	var usage float64

	for _, rate := range selectedPlan.Rates {
		if rate.Threshold != nil {
			// Calculates the charge for the rate
			charge := float64(*rate.Threshold) * rate.Price
			if total >= charge {
				// Calculates the remaining spend
				total -= charge
				// Adds the energy used
				usage += float64(*rate.Threshold)
			} else {
				// Adds the energy used
				usage += float64(*rate.Threshold)
				total = 0
			}
		} else {
			// Gets the remaining spend value divided by price and adds the usage so far
			usage = (float64(total) / rate.Price) + float64(usage)
		}
	}

	return &usage, nil
}

func addTax(value, taxRate float64) float64 {
	return value * float64(1+taxRate/100)
}

func removeTax(value, taxRate float64) float64 {
	return value / float64(1+taxRate/100)
}

func applyDiscount(total float64, plan data.EnergyPlan) float64 {
	newTotal := total

	for _, discount := range plan.Discount {
		switch discount.AppliesTo {
		case "whole_bill":
			newTotal = total - discount.Value
		}
	}

	return newTotal
}
