package new_product

import (
	"context"
	"fmt"

	"github.com/NekruzRakhimov/notification_service/internal/port/driven"
)

type UseCase struct {
	notifier          driven.Notifier
	authServiceClient driven.AuthServiceClient
}

func New(smtp driven.Notifier, authServiceClient driven.AuthServiceClient) *UseCase {
	return &UseCase{
		notifier:          smtp,
		authServiceClient: authServiceClient,
	}
}

func (uc *UseCase) Send(ctx context.Context, subject, body string) error {
	// Отправка запроса в auth_service на получение списка email'ов
	emails, err := uc.authServiceClient.GetAllEmails()
	if err != nil {
		return fmt.Errorf("usecase.Send: %w", err)
	}
	fmt.Printf("emails: %v\n", emails)

	// todo: важный момент - worker pool 50 - 25
	for _, email := range emails {
		if err := uc.notifier.Send(ctx, email, subject, body); err != nil {
			return err
		}
	}

	return nil
}
