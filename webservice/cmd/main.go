package main

import (
	"log"
	"net"
	"webservice/product"

	c "webservice"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	product.RegisterProductServiceServer(s, &c.API{})
	log.Println("running on port 8080")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
