package usecase

import (
	"github.com/NekruzRakhimov/notification_service/internal/adapter/driven/auth_service"
	"github.com/NekruzRakhimov/notification_service/internal/adapter/driven/smtp"
	"github.com/NekruzRakhimov/notification_service/internal/config"
	"github.com/NekruzRakhimov/notification_service/internal/port/usecase"
	"github.com/NekruzRakhimov/notification_service/internal/usecase/auth"
	"github.com/NekruzRakhimov/notification_service/internal/usecase/new_product"
)

type UseCases struct {
	Auth       usecase.Auth
	NewProduct usecase.NewProduct
}

func New(cfg config.Config) *UseCases {
	smtpNotifier := smtp.New(cfg.Smtp.User, cfg.Smtp.Password)
	authServiceClient := auth_service.New(cfg.AuthServiceURL)
	//so := simple_output.New()

	return &UseCases{
		Auth:       auth.New(smtpNotifier),
		NewProduct: new_product.New(smtpNotifier, authServiceClient),
	}
}
