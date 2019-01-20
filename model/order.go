package model

type Order struct {
	Name   string
	Tax    *Tax
	Price  float64
	Amount float64
}
