package main

import (
	"fmt"

	// "example.com/price-calculator/cmdmanager"
	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
	"example.com/price-calculator/utils"
)

const PRICE_FILE_PATH string = "data/prices.txt"

func main() {

	taxRate, err := utils.GetTaxRate()
	if err != nil {
		fmt.Println(err)
		return
	}
	//cmdm := cmdmanager.New() // use this var if u want input prices and get result from cmd
	fm := filemanager.New(PRICE_FILE_PATH, fmt.Sprintf("result_%0.f.json", taxRate*100))
	priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
	err = priceJob.Process()

	if err != nil {
		fmt.Println("Could not process job")
		fmt.Println(err)

	}

}
