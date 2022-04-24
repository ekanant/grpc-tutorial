package services

import (
	context "context"
	"fmt"

	"google.golang.org/grpc/status"
)

type calculationServer struct {
}

func NewCalculationServer() CalculatorServer {
	return &calculationServer{}
}

func (calculationServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	if req.Name == "" {
		return nil, status.Error(500, "Name is required")
	}
	result := &HelloResponse{
		Result: fmt.Sprintf("Hello %v", req.Name),
	}
	return result, nil
}
func (calculationServer) mustEmbedUnimplementedCalculatorServer() {}
