package custom_types

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Money struct {
	Amount   int64
	Currency string
}

func (a *Money) Scan(value interface{}) error {
	stringData, ok := value.(string)
	if !ok {
		fmt.Println("Asseertion error")
		return errors.New("Assertion error")
	}

	// Remove "()" from composite type
	dataString := strings.Trim(string(stringData), "()")
	values := strings.Split(dataString, ",")

	amount, err := strconv.ParseInt(values[0], 10, 64)
	if err != nil {
		return errors.New(err.Error())
	}

	currency := strings.Trim(values[1], "'")

	var res Money
	res = Money{
		Amount:   amount,
		Currency: currency,
	}

	*a = res
	return nil
}
