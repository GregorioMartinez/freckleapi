package freckleapi

type Invoice struct {
	ID          float64 `json:"id"`
	InvoiceDate string  `json:"invoice_date"`
	Reference   string  `json:"reference"`
	State       string  `json:"state"`
	TotalAmount float64 `json:"total_amount"`
	URL         string  `json:"url"`
}
