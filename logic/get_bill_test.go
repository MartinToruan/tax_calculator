package logic_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/MartinToruan/tax_calculator/logic"
	"github.com/MartinToruan/tax_calculator/persistence"
	. "github.com/smartystreets/goconvey/convey"
	sqlMock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestHandleGetBill(t *testing.T) {
	Convey("Success Handle Get Bill", t, func() {
		// Setup Persistence
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
			AddRow("Big Mac", 1, 1000)
		mock.ExpectQuery("select *").WillReturnRows(ret)

		// Construct Persistence
		persistenceClient := &persistence.TaxPersistence{
			MAX_DB_CLIENT: 1,
		}
		persistenceClient.SetClientPool(clients)
		defer persistenceClient.DeInit()

		// Setup Logic
		logicClient := &logic.TaxLogic{
			MAX_CONCURRENT: 1,
		}
		logicClient.Init()
		defer logicClient.DeInit()

		// Run Function
		bin, err := logicClient.HandleGetBill(persistenceClient)

		// Validate Function Return
		So(bin, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("Failed Handle Get Bill -- error when get data from database", t, func() {
		// Setup Persistence
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

		// Setup Logic
		logicClient := &logic.TaxLogic{
			MAX_CONCURRENT: 1,
		}
		logicClient.Init()
		defer logicClient.DeInit()

		// Run Function
		bin, err := logicClient.HandleGetBill(persistenceClient)

		// Validate Function Return
		So(bin, ShouldBeNil)
		So(err.Error(), ShouldEqual, "connection error")
	})
}
