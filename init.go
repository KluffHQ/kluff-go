package kluff

import (
	"context"
	"fmt"

	"github.com/KluffHQ/kluff-go/db"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func authInterceptor(token string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// TODO: Add authentication logic
		ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func Get(token string) (*Interactor, error) {
	conn, err := grpc.Dial("kluff_apps_core:9091",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(authInterceptor(token)),
	)
	if err != nil {
		return nil, fmt.Errorf("faild to connect to apps core: error: %v", err)
	}
	client := NewDBInteractor(conn)
	// Send Ping to the server to check if everything is working fine
	err = client.SendPing(context.Background(), &db.Ping{})
	if err != nil {
		return nil, err
	}
	return &client, nil
}
