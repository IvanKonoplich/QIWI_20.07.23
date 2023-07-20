package internal

import (
	"encoding/xml"
	"errors"
	"golang.org/x/net/html/charset"
	"net/http"
)

// GetAllRates получает все котировки
func GetAllRates(date, code string) ([]Valute, error) {

	client := &http.Client{}
	resp, err := client.Get("https://www.cbr-xml-daily.ru/daily.xml?date_req=" + date)
	if err != nil {
		return nil, err
	}
	var data Rates
	if resp.Status == "404 Not Found" {
		return nil, errors.New("the page was not found, a day off or an incorrect date may have been entered")
	}
	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}
	return data.V, nil
}
