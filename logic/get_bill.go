package logic

import (
	"encoding/json"
	"fmt"

	"github.com/MartinToruan/tax_calculator/model"
	"github.com/MartinToruan/tax_calculator/svc"
)

func (l *TaxLogic) HandleGetBill(p svc.Persistence) ([]byte, error) {
	// Get connection from poll
	<-l.connPool
	defer func() {
		// Insert back connection to poll
		l.connPool <- struct{}{}
	}()

	// Get All Data from Database
	items, err := p.GetAllData()
	if err != nil {
		fmt.Println("error while get data from database. err: ", err)
		return nil, err
	}

	// Construct billing
	billing := &model.Billing{
		Items: items,
	}
	billing.SetTotal()

	// Marshal Data
	data, _ := json.Marshal(billing)

	return data, nil
}
