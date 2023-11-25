package kluff

import (
	"context"

	"github.com/kluff-com/kluff-go/data/db"
	"github.com/kluff-com/kluff-go/data/helloworld"
	"github.com/kluff-com/kluff-go/pkg/internals"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func authInterceptor(token string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// TODO: Add authentication logic
		ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func New(token string) (*SDK, error) {
	conn, err := grpc.Dial("localhost:9091",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(authInterceptor(token)),
	)
	if err != nil {
		return nil, err
	}
	sdk := SDK{
		Interactor: internals.NewDBInteractor(conn),
		HelloWorld: helloworld.NewHelloWorldClient(conn),
	}
	// Send Ping to the server to check if everything is working fine
	err = sdk.SendPing(context.Background(), &db.Ping{})
	if err != nil {
		return nil, err
	}
	return &sdk, nil
}

type SDK struct {
	internals.Interactor
	HelloWorld helloworld.HelloWorldClient
}
