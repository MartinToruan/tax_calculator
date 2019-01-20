package model

import (
	"errors"
)

var mapCodeType = map[int]string{
	1: "Food & Beverage",
	2: "Tobacco",
	3: "Entertainment",
}

var mapCodeRefundable = map[int]bool{
	1: true,
	2: false,
	3: false,
}

var mapCodeAmountCalculation = map[int]func(price float64) float64{
	1: func(price float64) float64 { return float64(10/100.0) * price },
	2: func(price float64) float64 { return float64(2/100.0) * price },
	3: func(price float64) float64 { return float64(1/100.0) * (price - 100) },
}

type Tax struct {
	Code       int
	Type       string
	Refundable bool
	Amount     float64
}

// Define any Error Returned from Tax Object
var taxCodeNotFound = errors.New("tax code not found!")

// This function will Set TaxType based on TaxCode
// If the TaxType is not exist it will return Error
func (t *Tax) SetType() error {
	// Check Exist
	if taxType, ok := mapCodeType[t.Code]; ok {
		t.Type = taxType
		return nil
	}

	// Type not exist, return error
	return taxCodeNotFound
}

// This function will Set TaxRefundable based on TaxCode
// If TaxRefundable is not exist it will return Error
func (t *Tax) SetRefundable() error {
	// Check Exist
	if taxRef, ok := mapCodeRefundable[t.Code]; ok {
		t.Refundable = taxRef
		return nil
	}

	// Refundable not exist, return error
	return taxCodeNotFound
}

// This function will Set the TaAmount based on TaxCode and Price
// If the TaxCode is not exist it will return Error
func (t *Tax) SetAmount(price float64) error {
	// Check Exist
	if fn, ok := mapCodeAmountCalculation[t.Code]; ok {
		t.Amount = fn(price)
		return nil
	}

	// Amount Calculation func not exist, return err
	return taxCodeNotFound
}

// Function to Set mapCodeType
func SetTaxCodeType(data map[int]string) {
	mapCodeType = data
}

// Function to Set mapCodeRefundable
func SetTaxCodeRefundable(data map[int]bool) {
	mapCodeRefundable = data
}

// Function to Set mapCodeAmountCalculation
func SetTaxCodeAmountCalculation(data map[int]func(price float64) float64) {
	mapCodeAmountCalculation = data
}
