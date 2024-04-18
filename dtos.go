package main

type currency struct {
	CurrencyCode    string `xml:"CurrencyCode,attr"`
	BanknoteSelling string `xml:"BanknoteSelling"`
}

type currencies struct {
	Currency []currency
}
