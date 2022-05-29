package utils

import (
	"fmt"
	"net/smtp"
)

type SendEmailRequest struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Message   string `json:"message"`
}

type SendEmailResponse struct {
	Sent bool `json:"sent"`
}

type EmailService struct {
	config Config
}

type IEmailService interface {
	SendEmail(req SendEmailRequest) (resp SendEmailResponse, err error)
}

func (service *EmailService) SendEmail(req SendEmailRequest) (resp SendEmailResponse, err error) {
	auth := smtpAuthenticate(service.config)

	smtpAddr := fmt.Sprintf("%s:%d", service.config.SMTPHost, service.config.SMTPPort)

	err = smtp.SendMail(
		smtpAddr,
		auth,
		req.Recipient,
		[]string{req.Sender},
		formatMessage(req.Sender, req.Recipient, req.Subject, req.Message),
	)
	resp.Sent = err == nil

	return
}

func NewEmailService(config Config) IEmailService {
	return &EmailService{config: config}
}

/** Devuelve la autenticación del servicio SMTP
 *
 * @param config Config La configuración de la aplicación
 * @return smtp.Auth La autenticación
 */
func smtpAuthenticate(config Config) smtp.Auth {
	return smtp.PlainAuth(
		"",
		config.SMTPUser,
		config.SMTPPassword,
		config.SMTPHost,
	)
}

/** Devuelve una cadena de texto con el cuerpo del mensaje
 *
 * @param from string Remitente
 * @param to string Receptor
 * @param subject string Asunto del mensaje
 * @param message string Contenido del mensaje
 * @return []byte El cuerpo del mensaje
 */
func formatMessage(from, to, subject, body string) []byte {
	return []byte("From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		body + "\r\n")
}
