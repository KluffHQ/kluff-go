package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	api "github.com/kluff-com/kluff-go/proto"
)

func Add(x int32, y int32) int32 {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8085", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client for the Adder service.
	client := api.NewAdderClient(conn)

	// Prepare the request.
	req := &api.AddRequest{
		X: x,
		Y: y,
	}

	// Call the Add RPC.
	resp, err := client.Add(context.Background(), req)
	if err != nil {
		log.Fatalf("could not call Add RPC: %v", err)
	}
	return resp.Result
}

func main() {

	// Print the result.
	fmt.Printf("Result: %d\n", Add(3, 3))
}
