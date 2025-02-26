package utils

import (
	"errors"
	"fmt"
	"strconv"
)

func GetTaxRate() (float64, error) {

	var input string
	fmt.Println("Please enter tax rate, confirm with ENTER.")
	fmt.Print("Tax rate: ")
	fmt.Scan(&input)

	taxRate, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return taxRate, errors.New("could not convert input value to float")
	}

	if taxRate < 0 {
		return taxRate, errors.New("tax rate cannot be negative value")
	}

	if taxRate > 1 {
		return taxRate, errors.New("tax rate cannot be greater than 1")
	}

	return taxRate, nil
}

func StringsToFloat(inputs []string) ([]float64, error) {

	var result []float64

	for _, v := range inputs {

		floatPrice, err := strconv.ParseFloat(v, 64)

		if err != nil {

			return nil, errors.New("failed to convert string to float")

		}

		result = append(result, floatPrice)
	}

	return result, nil

}
