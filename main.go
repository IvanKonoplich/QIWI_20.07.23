package main

import (
	"QIWI_20.07.23/internal"
	"flag"
	"fmt"
	"log"
	"strings"
)

var codes = []string{
	"AUD",
	"GBP",
	"BYR",
	"DKK",
	"USD",
	"EUR",
	"ISK",
	"KZT",
	"CAD",
	"NOK",
	"XDR",
	"SGD",
	"TRL",
	"UAH",
	"SEK",
	"CHF",
	"JPY",
}

func main() {

	var (
		code string
		date string
	)
	//считываем флаги
	flag.StringVar(&code, "code", "", "code")
	flag.StringVar(&date, "date", "", "date")

	flag.Parse()

	if code == "" {
		log.Fatal("enter code")
	}
	if date == "" {
		log.Fatal("enter date")
	}

	if !checkCode(code) {
		log.Fatal("unknown code")
	}

	date = makeCorrectDate(date)
	rates, err := internal.GetAllRates(date, code)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	result, err := internal.FilterRates(code, rates)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	fmt.Println(result)
}

func makeCorrectDate(input string) string {
	inputSplit := strings.Split(input, "-")
	inputSplit[0], inputSplit[2] = inputSplit[2], inputSplit[0]
	result := strings.Join(inputSplit, "/")
	return result
}

func checkCode(c string) bool {
	for _, j := range codes {
		if j == c {
			return true
		}
	}
	return false
}
