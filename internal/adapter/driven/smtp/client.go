package smtp

import (
	"context"
	"fmt"

	"gopkg.in/gomail.v2"
)

type Client struct {
	Username string
	Password string
}

func New(username string, password string) *Client {
	return &Client{Username: username, Password: password}
}

const (
	GmailServerHost = "smtp.gmail.com"
	GmailServerPort = 587
)

func (c *Client) Send(ctx context.Context, recipient string, subject, body string) error {
	// Получаем credentials из переменных окружения

	if c.Username == "" || c.Password == "" {
		return fmt.Errorf("GMAIL credentials not set in environment variables")
	}

	// Создаём новое сообщение
	m := gomail.NewMessage()

	// Заголовки
	m.SetHeader("From", c.Username)
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", subject)

	// HTML тело письма
	htmlBody := c.buildHTMLBody(subject, body)
	m.SetBody("text/html", htmlBody)

	// Альтернативный plain text (для клиентов без HTML)
	m.AddAlternative("text/plain", body)

	// Настройка SMTP
	d := gomail.NewDialer("smtp.gmail.com", 587, c.Username, c.Password)

	// Отправка
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	fmt.Println("Email sent successfully to", recipient)
	return nil
}

func (c *Client) buildHTMLBody(subject, body string) string {
	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Your OTP Code</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f9f9f9;
            color: #333;
            line-height: 1.6;
            padding: 20px;
        }
        .container {
            max-width: 600px;
            margin: 0 auto;
            background-color: #fff;
            border-radius: 8px;
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
            padding: 20px;
        }
        .header {
            color: #4CAF50;
            text-align: center;
        }
        .otp-code {
            color: #FF5722;
            font-size: 28px;
            font-weight: bold;
            text-align: center;
            margin: 20px 0;
            padding: 15px;
            background-color: #f5f5f5;
            border-radius: 8px;
            letter-spacing: 5px;
        }
        .footer {
            font-size: 12px;
            color: #777;
            text-align: center;
            margin-top: 20px;
        }
    </style>
</head>
<body>
    <div class="container">
        <h2 class="header">%s</h2>
        <div class="otp-code">%s</div>
    </div>
</body>
</html>`, subject, body)
}
