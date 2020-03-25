package controllers

import (
	"fmt"
	"log"
	"os"
	"sort"

	"uswitch.com/comparison/calculations"
	"uswitch.com/comparison/data"
)

// ProcessPlanForCustomer -
func ProcessPlanForCustomer(calculator calculations.Calculator, plans []data.EnergyPlan, yearlyUsage int) {
	customerPlans := calculator.CalculatePlansForCustomer(plans, yearlyUsage)
	printPlans(customerPlans)
}

// ProcessEnergyUsedAnnually -
func ProcessEnergyUsedAnnually(calculator calculations.Calculator, plans []data.EnergyPlan, supplierName, planName string, monthlySpend int) {
	usage, err := calculator.CalculateEnergyUsedAnnually(plans, supplierName, planName, monthlySpend)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	// Using stdout to produce a similar output to the expected
	os.Stdout.WriteString(fmt.Sprintf("%.0f", *usage))
	os.Stdout.WriteString("\n")
}

func printPlans(customerPlans []data.CustomerPlan) {
	sort.Slice(customerPlans[:], func(i, j int) bool {
		return customerPlans[i].Total < customerPlans[j].Total
	})

	for _, cp := range customerPlans {
		// Using stdout to produce a similar output to the expected
		os.Stdout.WriteString(fmt.Sprintf("%s,%s,%s", cp.EnergyPlan.Supplier, cp.EnergyPlan.Plan, cp.TotalDisplay()))
		os.Stdout.WriteString("\n")
	}
}
