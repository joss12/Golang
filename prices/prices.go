package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/iomanager"
)

// var pl = fmt.Println

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InPutPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() error{

	lines, err := job.IOManager.ReadLines()
	if err != nil {

		return err
	}

	prices, err := conversion.StringToFloats(lines)
	if err != nil {
		return err
	}
	job.InPutPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process() error {
	err := job.LoadData()

	if err != nil{
		return err
	}

	result := make(map[string]string)


	for _, price := range job.InPutPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	job.TaxIncludedPrices = result
	return job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InPutPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
