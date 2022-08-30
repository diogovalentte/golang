package main

import (
	"context"
	"fmt"
	"log"

	"github.com/diogovalentte/golang/gRPC_server/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer connection.Close()

	client := pb.NewHelloServiceClient(connection)

	request := &pb.HelloRequest{
		Name: "Diogo",
	}

	res, err := client.Hello(context.Background(), request)
	if err != nil {
		log.Fatalf("Error during execution: %v", err)
	}
	fmt.Println(res)
}
