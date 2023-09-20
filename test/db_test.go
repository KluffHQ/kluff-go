package test

import (
	"context"
	"testing"

	"github.com/kluff-com/kluff-go"
)

func TestAuthentication(t *testing.T) {
	sdk, err := kluff.New(kluff.Config{
		APIKey: "some api Key",
	})
	if err != nil {
		t.Error(err)
	}

	data, err := sdk.GetFields(context.Background(), "hello")
	if err != nil {
		t.Error(err)
	}
	_ = data
}
