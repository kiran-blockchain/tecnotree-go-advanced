package calc

import "errors"

func Mul(a int, b int) int{
	return a*b
}
func Divide(a int,b int) (result int,err error){
	result=0
	err=nil
	if (b==0){
		err= errors.New("Divide by zero Error")
		//return result,errors.New("Divide by zero Error")
	}else{
		result= a/b
		//return result,nil
	}
	return
}