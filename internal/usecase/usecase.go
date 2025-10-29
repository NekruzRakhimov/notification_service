package usecase

import (
	"github.com/NekruzRakhimov/notification_service/internal/adapter/driven/smtp"
	"github.com/NekruzRakhimov/notification_service/internal/config"
	"github.com/NekruzRakhimov/notification_service/internal/port/usecase"
	"github.com/NekruzRakhimov/notification_service/internal/usecase/sender"
)

type UseCases struct {
	Sender usecase.Sender
}

func New(cfg config.Config) *UseCases {
	smtpNotifier := smtp.New(cfg.Smtp.User, cfg.Smtp.Password)
	//so := simple_output.New()

	return &UseCases{
		Sender: sender.New(smtpNotifier),
	}
}
