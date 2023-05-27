package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/go-resty/resty/v2"
)

type CurrencyData struct {
	Rates map[string]float64 `json:"rates"`
}

func ConvertCurrency(amount float64, targetCurrency string, rates map[string]float64) float64 {
	conversionRate := rates[targetCurrency]
	return amount * conversionRate
}

func main() {
	
	response, err := resty.New().R().Get("https://api.exchangerate-api.com/v4/latest/USD")
	if err != nil {
		fmt.Println("Failed to fetch exchange rates:", err)
		return
	}

	
	var data CurrencyData
	if err := json.Unmarshal(response.Body(), &data); err != nil {
		fmt.Println("Failed to parse response:", err)
		return
	}

	
	fmt.Println("Available currencies:")
	for currency := range data.Rates {
		fmt.Println(currency)
	}

	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the amount to convert: ")
	amountInput, _ := reader.ReadString('\n')
	amountInput = strings.TrimSpace(amountInput)
	amount, err := strconv.ParseFloat(amountInput, 64)
	if err != nil {
		fmt.Println("Invalid amount:", err)
		return
	}

	fmt.Print("Enter the target currency: ")
	targetCurrencyInput, _ := reader.ReadString('\n')
	targetCurrency := strings.TrimSpace(targetCurrencyInput)

	
	convertedAmount := ConvertCurrency(amount, targetCurrency, data.Rates)

	
	fmt.Printf("%.2f USD is equal to %.2f %s\n", amount, convertedAmount, targetCurrency)
}



// package main

// import (
// 	"encoding/json"
// 	"fmt"

// 	"github.com/go-resty/resty/v2"
// )

// type CurrencyData struct {
// 	Rates map[string]float64 `json:"rates"`
// }

// func ConvertCurrency(amount float64, targetCurrency string, rates map[string]float64) float64 {
// 	conversionRate := rates[targetCurrency]
// 	return amount * conversionRate
// }

// func main() {
// 	response, err := resty.New().R().Get("https://api.exchangerate-api.com/v4/latest/USD")
// 	if err != nil {
// 		fmt.Println("Failed to fetch exchange rates:", err)
// 		return
// 	}


// 	var data CurrencyData
// 	if err := json.Unmarshal(response.Body(), &data); err != nil {
// 		fmt.Println("Failed to parse response:", err)
// 		return
// 	}

// 	// Print available currencies
// 	fmt.Println("Available currencies:")
// 	for currency := range data.Rates {
// 		fmt.Println(currency)
// 	}


// 	usdAmount := 100.0
// 	targetCurrency := "EUR"
// 	convertedAmount := ConvertCurrency(usdAmount, targetCurrency, data.Rates)

// 	fmt.Printf("%.2f USD is equal to %.2f %s\n", usdAmount, convertedAmount, targetCurrency)
// }