package main

import (
	"context"
	"fmt"
	"os"

	pb "github.com/Sane4ek55/grpc_test/pkg"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	args := os.Args
	conn, err := grpc.Dial("127.0.0.1:5300", opts...)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := pb.NewReverseClient(conn)
	request := &pb.RequestReverse{
		Message: args[1],
	}
	response, err := client.Do(context.Background(), request)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	fmt.Println(response.Message)
}
