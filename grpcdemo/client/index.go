package main

import (
	"context"
	"fmt"
	"grpcdemo/profile"

	"google.golang.org/grpc"
)

func main(){
	conn,err := grpc.Dial(":6000",grpc.WithInsecure())
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer conn.Close()
	c:= profile.NewProfileServiceClient(conn)
	consumer := profile.Profile{Name:"Kiran",Email: "Kiranthecoder@gmail.com",Password: "1234"}
	result,err := c.CreateProfile(context.Background(),&consumer)
	if err!=nil{
		fmt.Println(err.Error())
	}else{
		fmt.Println(result)
	}

}