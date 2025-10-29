package simple_output

import (
	"context"
	"fmt"
)

type Client struct {
}

func New() *Client {
	return &Client{}
}

func (c *Client) Send(ctx context.Context, recipient string, subject, body string) error {
	fmt.Printf("Email sent successfully to %s\n", recipient)
	fmt.Printf("With subject %s\n", subject)
	fmt.Printf("With body %s\n", body)
	return nil
}
