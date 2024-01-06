package kluff

import (
	"github.com/kluff-com/kluff-go/sdk"
)

// get the kluff sdk instance.
// the token param is the should be the token the is parsed from the frontend
func Get(token string) (*SDK, error) {
	client, err := sdk.Get(token)
	if err != nil {
		return nil, err
	}
	return &SDK{
		Interactor: client,
	}, nil
}

type SDK struct {
	*sdk.Interactor
}
