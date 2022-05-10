package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"server/services"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	//startGrpc()
	startGrpcAndHttp()
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

func startHttp() {
	//grpcPort := "50051"
	httpPort := "3000"
	ctx := context.TODO()

	mux := runtime.NewServeMux()
	//opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := services.RegisterCalculatorHandlerServer(ctx, mux, services.NewCalculationServer()); err != nil {
		log.Fatalf("failed to start HTTP gateway: %v\n", err)
	}

	srv := &http.Server{
		Addr:    ":" + httpPort,
		Handler: mux,
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
		}

		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	log.Println("starting HTTP/REST gateway..." + httpPort)
	srv.ListenAndServe()
}
