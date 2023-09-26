// data_source.go
package stubbing

type DataSource interface {
    GetData() int
}

type MyDataSource struct{}

func (m *MyDataSource) GetData() int {
    return 42
}
