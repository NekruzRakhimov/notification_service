package bootstrap

import (
	"fmt"
	"github.com/NekruzRakhimov/notification_service/internal/adapter/driving/amqp"
	"github.com/NekruzRakhimov/notification_service/internal/config"
	"github.com/NekruzRakhimov/notification_service/internal/usecase"
)

func initLayers(cfg config.Config) *App {
	teardown := make([]func(), 0)
	//log := logger.New(cfg.LogLevel, config.ServiceLabel, zap.WithCaller(true))

	uc := usecase.New(cfg)

	//httpSrv := initHTTPService(&cfg, uc)
	amqpConn, amqpCh := amqp.InitAMQPConsumer(cfg.AMQPURL)

	teardown = append(teardown,
		func() {
			err := amqpCh.Close()
			if err != nil {
				fmt.Printf("Error closing AMQP consumer: %s\n", err)
				return
			}
		},
	)

	teardown = append(teardown,
		func() {
			err := amqpConn.Close()
			if err != nil {
				fmt.Printf("Error closing AMQP consumer: %s\n", err)
				return
			}
		},
	)

	productEventsQueue, err := amqp.InitProductEventsQueue(amqpCh)
	if err != nil {
		fmt.Printf("Error initializing product events queue: %s\n", err)
		return nil
	}

	authQueue, err := amqp.InitAuthQueue(amqpCh)
	if err != nil {
		fmt.Printf("Error initializing auth queue: %s\n", err)
		return nil
	}

	amqpConsumer := amqp.NewConsumersAMQP(productEventsQueue, authQueue, uc, amqpCh)

	return &App{
		cfg:          cfg,
		amqpConsumer: amqpConsumer,
		teardown:     teardown,
	}
}
