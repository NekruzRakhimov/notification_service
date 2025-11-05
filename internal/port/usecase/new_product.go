package usecase

import "context"

type NewProduct interface {
	Send(ctx context.Context, subject, body string) error
}
