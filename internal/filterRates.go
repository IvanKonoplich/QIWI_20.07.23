package internal

import (
	"errors"
	"fmt"
)

// FilterRates фильтрует полученные котировки
func FilterRates(code string, rates []Valute) (string, error) {
	for _, j := range rates {
		if j.CharCode == code {
			return fmt.Sprintf("%s (%s): %s", j.CharCode, j.Name, j.Value), nil
		}
	}
	return "", errors.New(fmt.Sprintf("no currency was received for - %s", code))
}
