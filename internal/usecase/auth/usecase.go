package auth

import (
	"context"

	"github.com/NekruzRakhimov/notification_service/internal/port/driven"
)

type UseCase struct {
	notifier driven.Notifier
}

func New(smtp driven.Notifier) *UseCase {
	return &UseCase{
		notifier: smtp,
	}
}

func (uc *UseCase) Send(ctx context.Context, recipient string) error {
	subject := "Ваш аккаунт успешно создан!"
	body := "Вы успешно зарегистрировались на платформе с самыми сочными ценами"
	if err := uc.notifier.Send(ctx, recipient, subject, body); err != nil {
		return err
	}

	return nil
}
