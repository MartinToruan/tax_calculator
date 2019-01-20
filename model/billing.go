package model

type Billing struct {
	Items         []*Item `json:"items"`
	PriceSubtotal float64 `json:"price_subtotal"`
	TaxSubTotal   float64 `json:"tax_subtotal"`
	GrandTotal    float64 `json:"grand_total"`
}

type Item struct {
	Name          string  `json:"name"`
	TaxCode       int     `json:"tax_code"`
	TaxType       string  `json:"type"`
	TaxRefundable string  `json:"refundable"`
	Price         float64 `json:"price"`
	TaxPrice      float64 `json:"tax"`
	Amount        float64 `json:"amount"`
}

func (b *Billing) SetTotal() {
	var pSubTotal, tSubTotal float64
	for _, item := range b.Items {
		pSubTotal += item.Price
		tSubTotal += item.TaxPrice
	}
	b.PriceSubtotal = pSubTotal
	b.TaxSubTotal = tSubTotal
	b.GrandTotal = pSubTotal + tSubTotal
}
