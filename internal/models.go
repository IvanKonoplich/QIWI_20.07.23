package internal

// Rates - это список всех котировок в представлении Valute
type Rates struct {
	V []Valute `xml:"Valute"`
}

// Valute - это представление котировки отдельной валюты
type Valute struct {
	CharCode string `xml:"CharCode"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}
