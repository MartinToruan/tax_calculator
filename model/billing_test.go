package model_test

import (
	"testing"

	"github.com/MartinToruan/tax_calculator/helper"
	"github.com/MartinToruan/tax_calculator/model"
	. "github.com/smartystreets/goconvey/convey"
)

func GetTestBilling() *model.Billing {
	ord1, _ := helper.ConstructOrder("Big Mac", 1, 1000)
	ord2, _ := helper.ConstructOrder("Lucky Stretch", 2, 1000)
	ord3, _ := helper.ConstructOrder("Movie", 3, 150)

	items := []*model.Item{
		helper.ConstructItemFromOrder(ord1),
		helper.ConstructItemFromOrder(ord2),
		helper.ConstructItemFromOrder(ord3),
	}

	return &model.Billing{
		Items: items,
	}
}

func TestSetTotal(t *testing.T) {
	Convey("Success Set Billing Total", t, func() {
		// Construct Billing Data Test
		billing := GetTestBilling()

		// Run Function
		billing.SetTotal()

		// Validate Function Return
		So(billing.PriceSubtotal, ShouldEqual, 2150)
		So(billing.TaxSubTotal, ShouldEqual, 120.5)
		So(billing.GrandTotal, ShouldEqual, 2270.5)
	})
}
