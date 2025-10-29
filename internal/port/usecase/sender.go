package usecase

import "context"

type Sender interface {
	Send(ctx context.Context, recipient string, subject, body string) error
}
