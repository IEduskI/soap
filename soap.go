package soap

import "context"

type Soap interface {
	Build(ctx context.Context, request Request) ([]byte, error)
	Send(ctx context.Context, request []byte) ([]byte, int, error)
}

type iflyRequest struct {
}

func NewIflyRequest() *iflyRequest {
	return &iflyRequest{}
}
