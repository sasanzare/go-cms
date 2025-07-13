package services

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/gomail.v2"
	"github.com/sasanzare/go-cms/utils"
)

// EmailService handles sending emails
type EmailService struct {
	dialer *gomail.Dialer
	sender string
}

// NewEmailService creates a new EmailService instance
func NewEmailService() *EmailService {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := 587 // default port, can be configured via env
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")
	sender := os.Getenv("EMAIL_SENDER")

	if smtpHost == "" || smtpUser == "" || smtpPass == "" {
		log.Println("Warning: SMTP configuration not fully set in environment variables")
	}

	return &EmailService{
		dialer: gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass),
		sender: sender,
	}
}

// EmailContent represents the content of an email
type EmailContent struct {
	To      string
	Subject string
	Body    string
	HTML    bool
}

// Send sends an email with the given content
func (es *EmailService) Send(content EmailContent) error {
	if !utils.ValidateEmail(content.To) {
		return fmt.Errorf("invalid recipient email address: %s", content.To)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", es.sender)
	m.SetHeader("To", content.To)
	m.SetHeader("Subject", content.Subject)

	if content.HTML {
		m.SetBody("text/html", content.Body)
	} else {
		m.SetBody("text/plain", content.Body)
	}

	if err := es.dialer.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}

// SendVerificationEmail sends an email with a verification link
func (es *EmailService) SendVerificationEmail(to, name, verificationURL string) error {
	subject := "Verify Your Email Address"
	body := fmt.Sprintf(`
Hello %s,

Please click the following link to verify your email address:
%s

If you didn't request this, please ignore this email.

Thanks,
The Team
`, name, verificationURL)

	return es.Send(EmailContent{
		To:      to,
		Subject: subject,
		Body:    strings.TrimSpace(body),
		HTML:    false,
	})
}

// SendPasswordResetEmail sends an email with a password reset link
func (es *EmailService) SendPasswordResetEmail(to, name, resetURL string) error {
	subject := "Password Reset Request"
	body := fmt.Sprintf(`
Hello %s,

We received a request to reset your password. Click the link below to proceed:
%s

This link will expire in 24 hours. If you didn't request a password reset, please ignore this email.

Thanks,
The Team
`, name, resetURL)

	return es.Send(EmailContent{
		To:      to,
		Subject: subject,
		Body:    strings.TrimSpace(body),
		HTML:    false,
	})
}

// SendWelcomeEmail sends a welcome email to new users
func (es *EmailService) SendWelcomeEmail(to, name string) error {
	subject := "Welcome to Our Platform!"
	body := fmt.Sprintf(`
Hello %s,

Welcome to our platform! We're excited to have you on board.

Here are some things you can do:
- Complete your profile
- Explore our features
- Contact support if you need help

Thanks for joining us!

Best regards,
The Team
`, name)

	return es.Send(EmailContent{
		To:      to,
		Subject: subject,
		Body:    strings.TrimSpace(body),
		HTML:    false,
	})
}