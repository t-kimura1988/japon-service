package main

import (
	"fmt"
	japonGrpc "github.com/t-kimura1988/japon-proto/pkg/grpc"
	"google.golang.org/grpc"
	//"jpon-service/src/types"
	"jpon-service/src/service/group"
	"log"
	"net"
)

func main() {
	port := 5200
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return
	}

	s := grpc.NewServer()

	japonGrpc.RegisterGroupServiceServer(s, group.NewGroupServer())

	if serverError := s.Serve(listener); serverError != nil {
		log.Fatalf("Failed to server: %v\n", serverError)
		return
	}
}
