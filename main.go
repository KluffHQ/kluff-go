package main

import (
	"github.com/kluff-com/kluff-go/data/helloworld"
	"github.com/kluff-com/kluff-go/pkg/internals"
	"google.golang.org/grpc"
)

type Config struct {
	APIKey string
}

func New(config Config) (kluffSDK, error) {
	conn, err := grpc.Dial("localhost:9091")
	if err != nil {
		return kluffSDK{}, err
	}

	sdk := kluffSDK{
		db:         internals.NewDBInteractor(conn),
		helloworld: helloworld.NewHelloWorldClient(conn),
	}

	return sdk, nil
}

type kluffSDK struct {
	db         internals.Interactor
	helloworld helloworld.HelloWorldClient
}
