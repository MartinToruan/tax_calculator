package svc

import (
	"database/sql"

	"github.com/MartinToruan/tax_calculator/model"
)

type Persistence interface {
	Init()
	DeInit()
	SetClientPool(clients chan *sql.DB)
	Insert(data *model.Order) (int, error)
	GetAllData() ([]*model.Item, error)
}
