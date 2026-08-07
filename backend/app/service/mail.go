package service

import (
	"fmt"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
	"rustdesk-api-server-pro/internal/errcode"
	"strings"
	"sync"

	mail "github.com/xhit/go-simple-mail/v2"
)

type MailService struct {
	mailer *mail.SMTPServer
	config *config.ServerConfig
}

var (
	mailServiceOnce sync.Once
	mailService     *MailService
)

func NewMailService() *MailService {
	mailServiceOnce.Do(func() {
		cfg := config.GetServerConfig()

		mailer := mail.NewSMTPClient()

		mailer.Host = cfg.SmtpConfig.Host
		mailer.Port = cfg.SmtpConfig.Port
		mailer.Username = cfg.SmtpConfig.Username
		mailer.Password = cfg.SmtpConfig.Password
		switch cfg.SmtpConfig.Encryption {
		case "ssl/tls":
			mailer.Encryption = mail.EncryptionSSLTLS
		case "starttls":
			mailer.Encryption = mail.EncryptionSTARTTLS
		default:
			mailer.Encryption = mail.EncryptionNone
		}

		mailService = &MailService{
			mailer: mailer,
			config: cfg,
		}
	})
	return mailService
}

func (service *MailService) GetMailTemplateByType(t int) (*model.MailTemplate, error) {
	var tpl model.MailTemplate
	has, err := db.DbEngine.Where("type = ?", t).Desc("id").Get(&tpl)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("mail template not found for type %d", t)
	}
	return &tpl, nil
}

func (service *MailService) Send(userId, tplId int, to, uuid string, vars map[string]string) error {

	sendLog := &model.MailLogs{
		UserId: userId,
		TplId:  tplId,
		From:   service.config.SmtpConfig.From,
		To:     to,
		Uuid:   uuid,
	}

	var template model.MailTemplate
	get, err := db.DbEngine.Where("id = ?", tplId).Get(&template)
	if err != nil || !get {
		sendLog.Status = model.MAIL_SEND_ERR
		errMsg := "mail template not found"
		if err != nil {
			errMsg = err.Error()
		}
		sendLog.Logs = errcode.Errorf(errcode.ERR8006.Code, errcode.ERR8006.Message+": "+errMsg).Error()
		db.DbEngine.Insert(sendLog)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s", errMsg)
	}

	body := template.Contents
	for k, v := range vars {
		body = strings.Replace(body, k, v, -1)
	}

	message := mail.NewMSG()
	message.SetFrom(service.config.SmtpConfig.From)
	message.AddTo(to)
	message.SetSubject(template.Subject)
	message.SetBody(mail.TextHTML, body)

	sender, err := service.mailer.Connect()
	if err != nil {
		sendLog.Status = model.MAIL_SEND_ERR
		sendLog.Logs = errcode.Errorf(errcode.ERR8007.Code, errcode.ERR8007.Message+": "+err.Error()).Error()
		db.DbEngine.Insert(sendLog)
		return err
	}
	err = message.Send(sender)
	if err != nil {
		sendLog.Status = model.MAIL_SEND_ERR
		sendLog.Logs = errcode.Errorf(errcode.ERR8008.Code, errcode.ERR8008.Message+": "+err.Error()).Error()
		db.DbEngine.Insert(sendLog)
		return err
	}

	sendLog.Subject = template.Subject
	sendLog.Contents = body
	sendLog.Status = model.MAIL_SEND_OK

	db.DbEngine.Insert(sendLog)

	return nil
}
