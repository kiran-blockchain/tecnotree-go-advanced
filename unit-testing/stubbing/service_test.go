// service_test.go
package stubbing

import (
    "testing"
)

type MockDataSource struct{}

func (m *MockDataSource) GetData() int {
    return 100
}

func TestMyService_Calculate(t *testing.T) {
    service := MyService{
        DataSource: &MockDataSource{},
    }

    result := service.Calculate()
    if result != 200 {
        t.Errorf("Expected 200, but got %d", result)
    }
}
