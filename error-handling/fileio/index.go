package fileio

import (
	"fmt"
	"log"
	"os"
)

func CreateFile(filename string, data string) {
	fmt.Println("Started Creating file \n")
	//step to create the file
	fileCreated,err:= os.Create(filename)
	if err!=nil{
		log.Fatalf("Failed to create the file",err)
	}
	//Close the file once the CreateFile completes
	defer fileCreated.Close()
	lengthOfData, err:= fileCreated.WriteString(data)
	 if err!=nil{
		log.Fatalf("Failed to write to the file",err)
	 }
	 fmt.Println(fileCreated.Name())
	 fmt.Println("Byte stored in the file",  lengthOfData)
}
func ReadFile(filename string){
	defer recoverReadingFile()	
	data,err:= os.ReadFile(filename)
	if err!=nil{
		//panic("Unable locate the file")
		log.Panicf("Failed to read the data from the file %s",err)
	}
	fmt.Printf("\n Data in the file:  %s",data)
}
func recoverReadingFile(){
	 r:=recover()//detect if the ReadFile function went into  panic
	 
	 if r!=nil{
		fmt.Println("I am recovered and will continue the flow")
	 }else{
		fmt.Println("There is no error happy path \n")
	 }

}
