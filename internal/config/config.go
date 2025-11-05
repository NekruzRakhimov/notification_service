package config

const ServiceLabel = "auth_service"

type Config struct {
	HTTPPort       string `env:"HTTP_PORT"`
	AMQPURL        string `env:"AMQP_URL"`
	Smtp           Smtp   `env:",prefix=SMTP_"`
	AuthServiceURL string `env:"AUTH_SERVICE_URL"`
}

type Smtp struct {
	User     string `env:"USER"`
	Password string `env:"PASSWORD" default:"owuv ciig uqyi ajxk"`
}
