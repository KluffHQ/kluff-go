package kluff

import (
	"context"

	"github.com/kluff-com/kluff-go/data/helloworld"
	"github.com/kluff-com/kluff-go/pkg/internals"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Config struct {
	APIKey string
}

func authInterceptor(Config Config) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// TODO: Add authtication logic
		ctx = metadata.AppendToOutgoingContext(ctx, "authorization", Config.APIKey)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func New(config Config) (kluffSDK, error) {
	conn, err := grpc.Dial("localhost:9091",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(authInterceptor(config)),
	)

	if err != nil {
		return kluffSDK{}, err
	}

	sdk := kluffSDK{
		Interactor: internals.NewDBInteractor(conn),
		HelloWorld: helloworld.NewHelloWorldClient(conn),
	}

	return sdk, nil
}

type kluffSDK struct {
	internals.Interactor
	HelloWorld helloworld.HelloWorldClient
}
