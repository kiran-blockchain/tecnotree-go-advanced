// add_test.go
package basics

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Expected 5, but got %d", result)
    }
}

func TestAddNegativeNumbers(t *testing.T) {
    result := Add(-2, -3)
    if result != -5 {
        t.Errorf("Expected -5, but got %d", result)
    }
}

func TestAddZero(t *testing.T) {
    result := Add(0, 0)
    if result != 0 {
        t.Errorf("Expected 0, but got %d", result)
    }
}
