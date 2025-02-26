package cmdmanager

import (
	"fmt"
)

type CMDManager struct {
}

func New() CMDManager {

	return CMDManager{}
}

func (cmdm CMDManager) ReadLines() ([]string, error) {

	var prices []string
	fmt.Println("Please enter your prices. Confirm every price with ENTER")
	fmt.Println("Enter 0 to process prices")
	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)
		if price == "0" {
			break
		}

		prices = append(prices, price)

	}

	return prices, nil
}

func (cmdm CMDManager) WriteResult(data interface{}) error {

	fmt.Println(data)

	return nil

}
