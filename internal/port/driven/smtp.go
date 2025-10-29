package driven

import "context"

type Notifier interface {
	Send(ctx context.Context, recipient string, subject, body string) error
}
