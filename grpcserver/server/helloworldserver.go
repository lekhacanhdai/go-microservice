package server

import (
	"context"
	"grpcserver/generated"
)

type HelloServer interface {
	Hello(ctx context.Context, request *generated.HelloRequest) (*generated.HelloResponse, error)
}
