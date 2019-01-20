package logic_test

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"
	"testing"

	"github.com/MartinToruan/tax_calculator/logic"
	"github.com/MartinToruan/tax_calculator/persistence"
	. "github.com/smartystreets/goconvey/convey"
	sqlMock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestHandleAddTax(t *testing.T) {
	Convey("Success Handle Add Tax", t, func() {
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
		ret := sqlMock.NewRows([]string{"id"}).AddRow(5)
		mock.ExpectQuery("insert into (.+)").WillReturnRows(ret)

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

		// Construct Request
		request_body := []byte(`
		{
			"name": "Big Mac",
			"tax_code": 1,
			"price": 1000
		}
		`)

		req, _ := http.NewRequest("POST", "/add/tax", bytes.NewBuffer(request_body))
		req.Header.Set("Content-Type", "application/json")

		// Run Function
		err = logicClient.HandleAddTax(req, persistenceClient)

		// Validate Function Return
		So(err, ShouldBeNil)
	})

	Convey("Failed Handle Add Tax -- decode error", t, func() {
		// Setup Persistence
		persistenceClient := &persistence.TaxPersistence{
			MAX_DB_CLIENT: 1,
		}

		// Setup Logic
		logicClient := &logic.TaxLogic{
			MAX_CONCURRENT: 1,
		}
		logicClient.Init()
		defer logicClient.DeInit()

		// Construct Request
		request_body := []byte(`
		{
			"name": 1,
			"tax_code": 1000,
			"price": "Big Mac"
		}
		`)

		req, _ := http.NewRequest("POST", "/add/tax", bytes.NewBuffer(request_body))
		req.Header.Set("Content-Type", "application/json")

		// Run Function
		err := logicClient.HandleAddTax(req, persistenceClient)

		// Validate Function Return
		So(err, ShouldNotBeNil)
	})

	Convey("Failed Handle Add Tax -- construct order error, invalid tax code", t, func() {
		// Setup Persistence
		persistenceClient := &persistence.TaxPersistence{
			MAX_DB_CLIENT: 1,
		}

		// Setup Logic
		logicClient := &logic.TaxLogic{
			MAX_CONCURRENT: 1,
		}
		logicClient.Init()
		defer logicClient.DeInit()

		// Construct Request
		request_body := []byte(`
		{
			"name": "Big Mac",
			"tax_code": 4,
			"price": 1000
		}
		`)

		req, _ := http.NewRequest("POST", "/add/tax", bytes.NewBuffer(request_body))
		req.Header.Set("Content-Type", "application/json")

		// Run Function
		err := logicClient.HandleAddTax(req, persistenceClient)

		// Validate Function Return
		So(err, ShouldNotBeNil)
	})

	Convey("Failed Handle Add Tax -- error while insert to database", t, func() {
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
		mock.ExpectQuery("insert into (.+)").WillReturnError(fmt.Errorf("connection error"))

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

		// Construct Request
		request_body := []byte(`
		{
			"name": "Big Mac",
			"tax_code": 1,
			"price": 1000
		}
		`)

		req, _ := http.NewRequest("POST", "/add/tax", bytes.NewBuffer(request_body))
		req.Header.Set("Content-Type", "application/json")

		// Run Function
		err = logicClient.HandleAddTax(req, persistenceClient)

		// Validate Function Return
		So(err.Error(), ShouldEqual, "connection error")
	})

}
