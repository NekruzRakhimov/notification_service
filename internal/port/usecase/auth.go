package usecase

import "context"

type Auth interface {
	Send(ctx context.Context, recipient string) error
}
