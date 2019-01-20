package svc

import "net/http"

type Logic interface {
	Init()
	DeInit()
	HandleAddTax(r *http.Request, p Persistence) error
	HandleGetBill(p Persistence) ([]byte, error)
}
