package closures


func Closure(client int) (func() int){
	//number := 10
	squareNum := func() int {
		client *= client
		return client
	}
	return squareNum
}
