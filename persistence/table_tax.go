package persistence

import (
	"github.com/MartinToruan/tax_calculator/helper"
	"github.com/MartinToruan/tax_calculator/model"
	_ "github.com/lib/pq"
)

// This function is used to Insert data to Database
// Return :
// 		- int -> row id
// 		- error -> if any
func (p *TaxPersistence) Insert(data *model.Order) (int, error) {
	// Get Database Client From Pool
	db := <-p.dbClientPool
	defer func() {
		// Insert back the Database Client to the Pool
		p.dbClientPool <- db
	}()

	// Prepare Query Insert
	sqlStatement := "insert into tax(name, tax_code, price) values($1, $2, $3) returning id"

	// Execute Query
	var lastInserted int
	if err := db.QueryRow(sqlStatement, data.Name, data.Tax.Code, data.Price).Scan(&lastInserted); err != nil {
		return 0, err
	}

	return lastInserted, nil
}

// This function is used to get all data from Database
// Return :
// 		- []*model.Item -> List of Data that already parsed to model.Item
//		- error 		-> if any
func (p *TaxPersistence) GetAllData() ([]*model.Item, error) {
	// Get Database Client From Pool
	db := <-p.dbClientPool
	defer func() {
		// Insert back the Database Client to the Pool
		p.dbClientPool <- db
	}()

	// Prepare Query Select
	sqlStatement := "select name, tax_code, price from tax"

	// Execute Query
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Loop Through Result
	var result []*model.Item
	for rows.Next() {
		// Scan each row. Each value will be mapped to related parameters.
		var name string
		var taxCode int
		var price float64
		err := rows.Scan(&name, &taxCode, &price)
		if err != nil {
			return nil, err
		}

		// Construct Order
		order, err := helper.ConstructOrder(name, taxCode, price)
		if err != nil {
			return nil, err
		}

		// Parse Order to item and insert to result
		result = append(result, helper.ConstructItemFromOrder(order))
	}
	return result, nil
}
