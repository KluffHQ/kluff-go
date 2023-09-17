package helloworld

import (
	"context"
	"log"

	"google.golang.org/grpc"

	helloWorldApi "github.com/kluff-com/kluff-go/pkg/api/helloworld"
)

type HelloWorldClient interface {
	HelloWorld(context.Context) (string, error)
}

type helloWorldClient struct {
	conn *grpc.ClientConn
}

func NewHelloWorldClient(address string) (HelloWorldClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &helloWorldClient{
		conn: conn,
	}, nil
}

func (c *helloWorldClient) HelloWorld(ctx context.Context) (string, error) {
	req := &helloWorldApi.HelloWorldRequest{}
	client := helloWorldApi.NewHelloWorldClient(c.conn)

	// Call the Add RPC.
	resp, err := client.HelloWorld(context.Background(), req)
	if err != nil {
		log.Fatalf("could not call HelloWorld RPC: %v", err)
		return "", err
	}
	return resp.Msg, nil
}

// This is for demo purpose to test grpc communication with apps-core repo
// This function will be used by apps backend as sdk.
func HelloWorldDemo() (string, error) {
	client, err := NewHelloWorldClient("localhost:9091")
	if err != nil {
		return "", err
	}

	return client.HelloWorld(context.Background())
}
