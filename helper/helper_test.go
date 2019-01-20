package helper_test

import (
	"testing"

	"github.com/MartinToruan/tax_calculator/helper"
	"github.com/MartinToruan/tax_calculator/model"
	. "github.com/smartystreets/goconvey/convey"
)

// Temp Valid CodeRefundable
var testCodeRefundable = map[int]bool{
	1: true,
	2: false,
	3: false,
}

// Temp Valid CodeAmountCalculation
var testCodeAmountCalculation = map[int]func(price float64) float64{
	1: func(price float64) float64 { return float64(10/100.0) * price },
	2: func(price float64) float64 { return float64(2/100.0) * price },
	3: func(price float64) float64 { return float64(1/100.0) * (price - 100) },
}

// Helper Function to generate Test data
func CreateTestRequestAddOrder() *model.RequestAddOrder {
	return &model.RequestAddOrder{
		Name:    "Big Mac",
		TaxCode: 1,
		Price:   1000,
	}
}

func CreateTestOrder(taxCode int) *model.Order {
	price := float64(1000)
	tax, _ := helper.ConstructTaxByTaxCode(taxCode, price)
	return &model.Order{
		Name:   "Big Mac",
		Tax:    tax,
		Price:  1000,
		Amount: price + tax.Amount,
	}
}

// Test Function
func TestConstructOrder(t *testing.T) {
	Convey("Success Construct Order", t, func() {
		// Construct Request Add Order Data
		req := CreateTestRequestAddOrder()

		// Run Function
		order, err := helper.ConstructOrder(req.Name, req.TaxCode, req.Price)

		// Validate Function Return
		So(err, ShouldBeNil)
		So(order.Name, ShouldEqual, "Big Mac")
		So(order.Tax.Code, ShouldEqual, 1)
		So(order.Tax.Type, ShouldEqual, "Food & Beverage")
		So(order.Tax.Refundable, ShouldEqual, true)
	})

	Convey("Failed Construct Order - Invalid Tax Type", t, func() {
		// Construct Request Add Order Data
		req := CreateTestRequestAddOrder()

		// Set Invalid Value
		req.TaxCode = 10

		// Run Function
		order, err := helper.ConstructOrder(req.Name, req.TaxCode, req.Price)

		// Validate Function Return
		So(err, ShouldNotBeNil)
		So(order, ShouldBeNil)
	})

	Convey("Failed Construct Order - Invalid Refundable", t, func() {
		// Construct Request Add Order Data
		req := CreateTestRequestAddOrder()

		// Set Invalid mapCodeRefundable
		model.SetTaxCodeRefundable(nil)
		defer model.SetTaxCodeRefundable(testCodeRefundable)

		// Run Function
		order, err := helper.ConstructOrder(req.Name, req.TaxCode, req.Price)

		// Validate Function Return
		So(order, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})
}
func TestConstructTaxByTaxCode(t *testing.T) {
	Convey("Success Construct Tax By TaxCode", t, func() {
		// Run Function
		tax, err := helper.ConstructTaxByTaxCode(1, 1000)

		// Validate Function Return
		So(err, ShouldBeNil)
		So(tax.Code, ShouldEqual, 1)
		So(tax.Type, ShouldEqual, "Food & Beverage")
		So(tax.Refundable, ShouldEqual, true)
		So(tax.Amount, ShouldEqual, float64(10/100.0)*1000)
	})

	Convey("Failed Construct Tax By TaxCode - Invalid Tax Type", t, func() {
		// Run Function
		tax, err := helper.ConstructTaxByTaxCode(10, 1000)

		// Validate Function Return
		So(tax, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})

	Convey("Failed Construct Tax By TaxCode - Invalid Refundable", t, func() {
		// Set Invalid mapCodeRefundable
		model.SetTaxCodeRefundable(nil)
		defer model.SetTaxCodeRefundable(testCodeRefundable)

		// Run Function
		tax, err := helper.ConstructTaxByTaxCode(1, 1000)

		// Validate Function Return
		So(tax, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})

	Convey("Failed Construct Tax By TaxCode - Invalid Amount", t, func() {
		// Set Invalid mapCodeRefundable
		model.SetTaxCodeAmountCalculation(nil)
		defer model.SetTaxCodeAmountCalculation(testCodeAmountCalculation)

		// Run Function
		tax, err := helper.ConstructTaxByTaxCode(1, 1000)

		// Validate Function Return
		So(tax, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})
}

func TestConstructItemFromOrder(t *testing.T) {
	Convey("Success Construct Item From Order - Refundable true", t, func() {
		// Construct Data Test
		order := CreateTestOrder(1)

		// Run Function
		item := helper.ConstructItemFromOrder(order)

		// Validate Function Return
		So(item.Name, ShouldEqual, order.Name)
		So(item.TaxCode, ShouldEqual, order.Tax.Code)
		So(item.TaxType, ShouldEqual, order.Tax.Type)
		So(item.Price, ShouldEqual, order.Price)
		So(item.TaxPrice, ShouldEqual, order.Tax.Amount)
		So(item.Amount, ShouldEqual, order.Amount)
		So(item.TaxRefundable, ShouldEqual, "Yes")
	})

	Convey("Success Construct Item From Order - Refundable false", t, func() {
		// Construct Data Test
		order := CreateTestOrder(2)

		// Run Function
		item := helper.ConstructItemFromOrder(order)

		// Validate Function Return
		So(item.Name, ShouldEqual, order.Name)
		So(item.TaxCode, ShouldEqual, order.Tax.Code)
		So(item.TaxType, ShouldEqual, order.Tax.Type)
		So(item.Price, ShouldEqual, order.Price)
		So(item.TaxPrice, ShouldEqual, order.Tax.Amount)
		So(item.Amount, ShouldEqual, order.Amount)
		So(item.TaxRefundable, ShouldEqual, "No")
	})
}
