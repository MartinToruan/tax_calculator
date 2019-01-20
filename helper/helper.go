package helper

import (
	"fmt"

	"github.com/MartinToruan/tax_calculator/model"
)

// Usage:
// Construct Order Model from RequestAddOrder Model.
func ConstructOrder(name string, taxCode int, price float64) (*model.Order, error) {
	// Construct Tax
	tax, err := ConstructTaxByTaxCode(taxCode, price)
	if err != nil {
		return nil, err
	}

	// Construct Order
	order := &model.Order{
		Name:   name,
		Tax:    tax,
		Price:  price,
		Amount: price + tax.Amount,
	}

	return order, nil
}

func ConstructTaxByTaxCode(taxCode int, price float64) (*model.Tax, error) {
	// Construct Tax Model
	tax := &model.Tax{
		Code: taxCode,
	}

	// Set Type
	if err := tax.SetType(); err != nil {
		fmt.Println("error when set Tax Type. error: ", err)
		return nil, err
	}

	// Set Refundable
	if err := tax.SetRefundable(); err != nil {
		fmt.Println("error when set Tax Refundable. error: ", err)
		return nil, err
	}

	// Set Amount
	if err := tax.SetAmount(price); err != nil {
		fmt.Println("error when set Tax Amount. error: ", err)
		return nil, err
	}

	return tax, nil
}

func ConstructItemFromOrder(data *model.Order) *model.Item {
	// Construct Item
	item := &model.Item{
		Name:     data.Name,
		TaxCode:  data.Tax.Code,
		TaxType:  data.Tax.Type,
		Price:    data.Price,
		TaxPrice: data.Tax.Amount,
		Amount:   data.Amount,
	}

	// Change Refundable Value to Sentence "Yes" or "No"
	if data.Tax.Refundable {
		item.TaxRefundable = "Yes"
	} else {
		item.TaxRefundable = "No"
	}

	return item
}
