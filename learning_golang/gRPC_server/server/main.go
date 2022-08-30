package main

import (
	"context"
	"errors"
	"log"
	"net"

	"github.com/JohnCGriffin/overflow"
	"github.com/diogovalentte/golang/gRPC_server/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.CheckStatusServiceServer
	pb.CalculateServiceServer
}

func (*server) CheckStatus(ctx context.Context, req *pb.CheckStatusRequest) (*pb.CheckStatusResponse, error) {
	response := &pb.CheckStatusResponse{
		Code: 200,
	}
	return response, nil
}

func (*server) Calculate(ctx context.Context, req *pb.CalculateRequest) (*pb.CalculateResponse, error) {
	num1 := req.Num1
	num2 := req.Num2
	operator := req.Operator

	// Do math depending on the arithmetic operator
	var result int64
	var responseErr error = nil

	switch operator {
	case "+":
		r, ok := overflow.Add64(num1, num2) // Check if num1 + num2 dont overflow
		if !ok {
			errMsg := "the result of Number 1 + Number 2 overflow"
			responseErr = errors.New(errMsg)
		} else {
			result = r
		}
	case "-":
		r, ok := overflow.Sub64(num1, num2) // Check if num1 - num2 dont overflow
		if !ok {
			errMsg := "the result of Number 1 - Number 2 overflow"
			responseErr = errors.New(errMsg)
		} else {
			result = r
		}
	case "*":
		r, ok := overflow.Mul64(num1, num2) // Check if num1 * num2 dont overflow
		if !ok {
			errMsg := "the result of Number 1 * Number 2 overflow"
			responseErr = errors.New(errMsg)
		} else {
			result = r
		}
	case "/":
		r, ok := overflow.Div64(num1, num2) // Check if num1 / num2 dont overflow
		if !ok {
			errMsg := "the result of Number 1 / Number 2 overflow"
			responseErr = errors.New(errMsg)
		} else {
			result = r
		}
	}

	response := &pb.CalculateResponse{
		Num1:     num1,
		Num2:     num2,
		Operator: operator,
		Result:   result,
	}

	return response, responseErr
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCheckStatusServiceServer(grpcServer, &server{})
	pb.RegisterCalculateServiceServer(grpcServer, &server{})

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
