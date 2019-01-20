package model_test

import (
	"testing"

	"github.com/MartinToruan/tax_calculator/model"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSetType(t *testing.T) {
	Convey("Success Set Type", t, func() {
		// Construct Tax Data Test
		tax := &model.Tax{Code: 1}

		// Run Function
		err := tax.SetType()

		// Validate Function Return
		So(err, ShouldBeNil)
		So(tax.Type, ShouldEqual, "Food & Beverage")
	})

	Convey("Failed Set Type -- tax code not exist", t, func() {
		// Construct Tax Data Test
		tax := &model.Tax{Code: 4}

		// Run Function
		err := tax.SetType()

		// Validate Function Return
		So(err.Error(), ShouldEqual, "tax code not found!")
	})
}

func TestSetRefundable(t *testing.T) {
	Convey("Success Set Refundable", t, func() {
		// Construct Tax Data Test
		tax := &model.Tax{Code: 1}

		// Run Function
		err := tax.SetRefundable()

		// Validate Function Return
		So(err, ShouldBeNil)
		So(tax.Refundable, ShouldEqual, true)
	})

	Convey("Failed Set Refundable -- tax code not exist", t, func() {
		// Construct Tax Data Test
		tax := &model.Tax{Code: 4}

		// Run Function
		err := tax.SetRefundable()

		// Validate Function Return
		So(err.Error(), ShouldEqual, "tax code not found!")
	})
}

func TestSetAmount(t *testing.T) {
	Convey("Success Set Amount", t, func() {
		// Construct Tax Data Test
		tax := &model.Tax{Code: 1}

		// Run Function
		err := tax.SetAmount(1000)

		// Validate Function Return
		So(err, ShouldBeNil)
		So(tax.Amount, ShouldEqual, 100)
	})

	Convey("Failed Set Amount", t, func() {
		// Construct Tax Data Test
		tax := &model.Tax{Code: 4}

		// Run Function
		err := tax.SetAmount(1000)

		// Validate Function Return
		So(err.Error(), ShouldEqual, "tax code not found!")
	})
}
