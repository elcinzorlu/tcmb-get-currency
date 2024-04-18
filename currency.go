package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Currency string

const (
	USD Currency = "USD"
	EUR Currency = "EUR"
	GBP Currency = "GBP"
)

func GetTCMBXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}

func GetCurrencies() map[Currency]string {
	var result currencies
	currencyMap := make(map[Currency]string, 0)
	if xmlBytes, err := GetTCMBXML("https://www.tcmb.gov.tr/kurlar/today.xml"); err != nil {
		fmt.Printf("Failed to get XML: %v", err)
	} else {
		xml.Unmarshal(xmlBytes, &result)
		currencyArr := []Currency{USD, EUR, GBP}
		for _, c := range result.Currency {
			for _, currency := range currencyArr {
				if c.CurrencyCode == string(currency) {
					currencyMap[currency] = c.BanknoteSelling
				}
			}
		}
	}
	return currencyMap
}

func ConvertCurrency(amount float64, curr Currency) float64 {
	getCurr := GetCurrencies()
	currPrice := getCurr[curr]
	var calculateCurrPrice float64
	if amount != 0 {
		fCurrPrice, err := strconv.ParseFloat(currPrice, 64)
		if err != nil {
			fmt.Println("Error:", err)
			return 0
		}
		calculateCurrPrice = amount / fCurrPrice
	}
	return roundFloat(calculateCurrPrice, 4)

}
