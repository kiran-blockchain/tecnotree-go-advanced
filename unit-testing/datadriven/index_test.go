package datadriven
import (
	"testing"
)

func TestCategorizeInvoices(t *testing.T) {
	testCases := []struct {
		input    []Invoice
		expected InvoiceCategory
	}{
		{
			[]Invoice{
				{1, 50.0},
				{2, 75.0},
				{3, 150.0},
				{4, 900.0},
				{5, 1100.0},
			},
			InvoiceCategory{
				Low:    []Invoice{{1, 50.0}, {2, 75.0}},
				Medium: []Invoice{{3, 150.0}, {4, 900.0}},
				High:   []Invoice{{5, 1100.0}},
			},
		},
		{
			[]Invoice{
				{6, 200.0},
				{7, 800.0},
			},
			InvoiceCategory{
				Low:    []Invoice{},
				Medium: []Invoice{{6, 200.0}},
				High:   []Invoice{{7, 800.0}},//negative case
			},
		},
		{
			[]Invoice{},
			InvoiceCategory{
				Low:    []Invoice{},
				Medium: []Invoice{},
				High:   []Invoice{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run("Testing invoice categorization", func(t *testing.T) {
			result := CategorizeInvoices(tc.input)

			// Check Low category
			if !equalInvoices(result.Low, tc.expected.Low) {
				t.Errorf("Low category mismatch. Expected: %+v, Got: %+v", tc.expected.Low, result.Low)
			}

			// Check Medium category
			if !equalInvoices(result.Medium, tc.expected.Medium) {
				t.Errorf("Medium category mismatch. Expected: %+v, Got: %+v", tc.expected.Medium, result.Medium)
			}

			// Check High category
			if !equalInvoices(result.High, tc.expected.High) {
				t.Errorf("High category mismatch. Expected: %+v, Got: %+v", tc.expected.High, result.High)
			}
		})
	}
}

func equalInvoices(a, b []Invoice) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
