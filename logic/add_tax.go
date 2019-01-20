package logic

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MartinToruan/tax_calculator/helper"
	"github.com/MartinToruan/tax_calculator/model"
	"github.com/MartinToruan/tax_calculator/svc"
)

func (l *TaxLogic) HandleAddTax(r *http.Request, p svc.Persistence) error {
	// Get connection from poll
	<-l.connPool
	defer func() {
		// Insert back connection to poll
		l.connPool <- struct{}{}
	}()

	// Decode Message Body
	var reqData model.RequestAddOrder
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		fmt.Println("error when decode message. err: ", err)
		return err
	}

	order, err := helper.ConstructOrder(reqData.Name, reqData.TaxCode, reqData.Price)
	if err != nil {
		fmt.Println("error while construct order from request. err: ", err)
		return err
	}

	lastInsertedId, err := p.Insert(order)
	if err != nil {
		fmt.Println("error while inserted to database. err: ", err)
		return err
	}
	fmt.Println("success insert data to database. ID: ", lastInsertedId)

	return nil
}
