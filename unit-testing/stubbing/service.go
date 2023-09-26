package stubbing

type MyService struct {
    DataSource DataSource
}

func (s *MyService) Calculate() int {
    return s.DataSource.GetData() * 2
}