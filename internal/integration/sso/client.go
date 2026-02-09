package sso

import "context"

type Identify struct {
	ID    int64
	Login string
}

type Client interface {
	Me(ctx context.Context, token string) (*Identify, error)
}
