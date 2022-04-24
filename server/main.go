package main

import (
	"log"
	"net"
	"server/services"

	"google.golang.org/grpc"
)

func main() {
	startGrpc()
}

func startGrpc() {
	s := grpc.NewServer()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	services.RegisterCalculatorServer(s, services.NewCalculationServer())
	log.Println("Serve grpc service")
	err = s.Serve(listener)
	if err != nil {
		log.Fatalln(err)
	}
}
