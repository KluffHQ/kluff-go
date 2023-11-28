package kluff

import (
	"context"

	"github.com/kluff-com/kluff-go/data/account"
	"github.com/kluff-com/kluff-go/data/db"
	"github.com/kluff-com/kluff-go/pkg/internals"
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

// get the kluff sdk instance.
// the token param is the should be the token the is parsed from the frontend
func Get(token string) (*SDK, error) {
	conn, err := grpc.Dial("localhost:9091",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(authInterceptor(token)),
	)
	if err != nil {
		return nil, err
	}
	sdk := SDK{
		Interactor: internals.NewDBInteractor(conn),
		Accounting: account.NewAccountClient(conn),
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
	Accounting account.AccountClient
}
