package main

import (
	"context"
	"fmt"
	"github.com/Jayleonc/awesome/grpc/myservice"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := myservice.NewMyServiceClient(conn)

	message := "World"
	fmt.Println(message)
	resp, err := c.MyMethod(context.Background(), &myservice.MyRequest{Message: message})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response: %s", resp.Message)
}
