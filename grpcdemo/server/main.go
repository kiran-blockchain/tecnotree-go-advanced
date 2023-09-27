package main

import (
	"context"
	"fmt"
	"grpcdemo/profile"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	profile.UnimplementedProfileServiceServer
}

func (*server) CreateProfile(ctx context.Context,
	 request *profile.Profile) (*profile.ProfileResponse, error) {
	return &profile.ProfileResponse{Id: "1",
	 ErrorMessage: "", Success: request.Email + " Successfully Created"}, nil
}
func main() {
	lis, err := net.Listen("tcp", ":6000")
	defer lis.Close()
	if err != nil {
		log.Printf(err.Error())
		return
	}
	s := grpc.NewServer()
	profile.RegisterProfileServiceServer(s, &server{})
	fmt.Println("Server listening on port :6000")
	if err := s.Serve(lis); err != nil {
		fmt.Print(err.Error())
	}
}
