package prices

import (
	"fmt"

	"example.com/price-calculator/iomanager"
	"example.com/price-calculator/utils"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]float64  `json:"tax_included_prices"`
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		IOManager: iom,
		TaxRate:   taxRate,
	}
}

func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()
	if err != nil {

		return err
	}

	prices, err := utils.StringsToFloat(lines)

	if err != nil {

		return err

	}

	job.InputPrices = prices

	return nil

}

func (job *TaxIncludedPriceJob) Process() error {
	job.LoadData()
	result := make(map[string]float64)

	for _, price := range job.InputPrices {

		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)

	}
	job.TaxIncludedPrices = result

	return job.IOManager.WriteResult(job)
}
