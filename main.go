package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/alecthomas/kingpin"
	"uswitch.com/energy-comparison/calculations"
	"uswitch.com/energy-comparison/controllers"
	"uswitch.com/energy-comparison/data"
)

var (
	plansFile = kingpin.Arg("plans-path", "Path for the file with plans").Required().String()
	taxRate   = kingpin.Flag("tax-rate", "Tax rate percentage for energy prices").Default("5").Float()
)

func main() {
	kingpin.Parse()

	plansData, err := ioutil.ReadFile(*plansFile)
	if err != nil {
		log.Fatal(err)
	}

	plans := make([]data.EnergyPlan, 0)

	if err := json.Unmarshal(plansData, &plans); err != nil {
		log.Fatal(err)
	}

	defaultCaltulator := calculations.NewDefaultCalculator(*taxRate)

	reader := bufio.NewReader(os.Stdin)
	command := ""

	for command != "exit" {

		command, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		command = strings.TrimSuffix(command, "\n")

		args := strings.Split(command, " ")

		if len(args) > 0 {
			switch args[0] {

			case "price":
				usage := getPriceArgs(args)
				controllers.ProcessPlanForCustomer(defaultCaltulator, plans, usage)

			case "usage":
				supplierName, planName, monthlySpend := getUsageArgs(args)
				controllers.ProcessEnergyUsedAnnually(defaultCaltulator, plans, supplierName, planName, monthlySpend)

			case "exit":

			default:
				printHelp()
			}
		}

	}
}

func getPriceArgs(args []string) int {
	if len(args) != 2 {
		log.Fatal("Error : price command requires a monthly input. (e.g price 1000)")
	}

	usage, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return int(usage)
}

func getUsageArgs(args []string) (string, string, float64) {

	//if len(args) != 4 {
	//log.Fatal("Error : usage command requires supplier-name plan-name monthly-spend. (e.g usage xxx yyy 100)")
	//}

	supplierName := args[1]
	planName := strings.Join(args[2:len(args)-1], " ")
	monthlySpend, err := strconv.ParseFloat(args[len(args)-1], 64)
	if err != nil {
		log.Fatal(err)
	}

	return supplierName, planName, monthlySpend
}

func printHelp() {
	log.Printf("------- HELP -------")
	log.Printf("price command calculates the plans for a given customer. It requires a monthly input. (e.g price 1000)")
	log.Printf("usage command calculates the amount of energy used with a given provider. It requires supplier-name plan-name monthly-spend. (e.g usage xxx yyy 100)")
	log.Printf("exit command leaves the program")
	log.Printf("--------------------")
}
