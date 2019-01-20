package persistence_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/MartinToruan/tax_calculator/helper"
	"github.com/MartinToruan/tax_calculator/model"
	"github.com/MartinToruan/tax_calculator/persistence"
	. "github.com/smartystreets/goconvey/convey"
	sqlMock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

// Helper Function
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

func TestInsert(t *testing.T) {
	Convey("Success Insert", t, func() {
		// Generate Sql Mock
		clients := make(chan *sql.DB, 1)

		db, mock, err := sqlMock.New()
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		clients <- db

		// Construct Return
		ret := sqlMock.NewRows([]string{"id"}).AddRow(5)
		mock.ExpectQuery("insert into (.+)").WillReturnRows(ret)

		// Construct Persistence
		persistenceClient := &persistence.TaxPersistence{
			MAX_DB_CLIENT: 1,
		}
		persistenceClient.SetClientPool(clients)
		defer persistenceClient.DeInit()

		// Construct Data Test
		order := CreateTestOrder(1)

		// Run Function
		lastInsertedId, err := persistenceClient.Insert(order)

		// Validate Function Return
		So(err, ShouldBeNil)
		So(lastInsertedId, ShouldEqual, 5)
	})

	Convey("Failed Insert -- connection error", t, func() {
		// Generate Sql Mock
		clients := make(chan *sql.DB, 1)

		db, mock, err := sqlMock.New()
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		clients <- db

		// Construct Return
		mock.ExpectQuery("insert into (.+)").WillReturnError(fmt.Errorf("connection error"))

		// Construct Persistence
		persistenceClient := &persistence.TaxPersistence{
			MAX_DB_CLIENT: 1,
		}
		persistenceClient.SetClientPool(clients)
		defer persistenceClient.DeInit()

		// Construct Data Test
		order := CreateTestOrder(1)

		// Run Function
		lastInsertedId, err := persistenceClient.Insert(order)

		// Validate Function Return
		So(err.Error(), ShouldEqual, "connection error")
		So(lastInsertedId, ShouldEqual, 0)
	})
}

func TestGetAllData(t *testing.T) {
	Convey("Success Get AllData", t, func() {
		// Generate Sql Mock
		clients := make(chan *sql.DB, 1)

		db, mock, err := sqlMock.New()
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		clients <- db

		// Construct Return
		ret := sqlMock.NewRows([]string{"name", "tax_code", "price"}).
			AddRow("Big Mac", 1, 1000).
			AddRow("Lucky Stretch", 2, 1000)
		mock.ExpectQuery("select *").WillReturnRows(ret)

		// Construct Persistence
		persistenceClient := &persistence.TaxPersistence{
			MAX_DB_CLIENT: 1,
		}
		persistenceClient.SetClientPool(clients)
		defer persistenceClient.DeInit()

		// Run Function
		items, err := persistenceClient.GetAllData()

		// Validate Function Return
		So(err, ShouldBeNil)
		So(len(items), ShouldEqual, 2)

		// Item 1
		So(items[0].Name, ShouldEqual, "Big Mac")
		So(items[0].TaxCode, ShouldEqual, 1)
		So(items[0].TaxType, ShouldEqual, "Food & Beverage")
		So(items[0].TaxRefundable, ShouldEqual, "Yes")
		So(items[0].Price, ShouldEqual, 1000)
		So(items[0].TaxPrice, ShouldEqual, 100)
		So(items[0].Amount, ShouldEqual, 1100)

		// Item 2
		So(items[1].Name, ShouldEqual, "Lucky Stretch")
		So(items[1].TaxCode, ShouldEqual, 2)
		So(items[1].TaxType, ShouldEqual, "Tobacco")
		So(items[1].TaxRefundable, ShouldEqual, "No")
		So(items[1].Price, ShouldEqual, 1000)
		So(items[1].TaxPrice, ShouldEqual, 20)
		So(items[1].Amount, ShouldEqual, 1020)
	})

	Convey("Failed Get AllData -- connection error", t, func() {
		// Generate Sql Mock
		clients := make(chan *sql.DB, 1)

		db, mock, err := sqlMock.New()
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
			return
		}
		clients <- db

		// Construct Return
		mock.ExpectQuery("select *").WillReturnError(fmt.Errorf("connection error"))

		// Construct Persistence
		persistenceClient := &persistence.TaxPersistence{
			MAX_DB_CLIENT: 1,
		}
		persistenceClient.SetClientPool(clients)
		defer persistenceClient.DeInit()

		// Run Function
		items, err := persistenceClient.GetAllData()

		// Validate Function Return
		So(items, ShouldBeNil)
		So(err.Error(), ShouldEqual, "connection error")
	})
}
