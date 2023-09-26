package datadriven


type Invoice struct {
	ID    int
	Value float64
}

type InvoiceCategory struct {
	Low    []Invoice
	Medium []Invoice
	High   []Invoice
}

func CategorizeInvoices(invoices []Invoice) InvoiceCategory {
	var result InvoiceCategory

	for _, invoice := range invoices {
		if invoice.Value < 100 {
			result.Low = append(result.Low, invoice)
		} else if invoice.Value < 1000 {
			result.Medium = append(result.Medium, invoice)
		} else {
			result.High = append(result.High, invoice)
		}
	}

	return result
}


