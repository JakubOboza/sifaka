package notifications

import (
	"bytes"
	"fmt"

	"github.com/JakubOboza/sifaka/config"
	"github.com/JakubOboza/sifaka/storage"
	"github.com/slack-go/slack"
)

type SlackService struct {
	oauthToken string
	channelID  string
}

func NewSlackProvider() Provider {
	oauthToken := config.SlackOAuth()
	channelID := config.SlackChannelID()
	return &SlackService{oauthToken: oauthToken, channelID: channelID}
}

func (svc *SlackService) Notify(certs []storage.CertificateData) error {

	attachment := slack.Attachment{
		Text: certWarningMessageText(certs),
	}

	return svc.sendMessage(attachment)
}

func (svc *SlackService) Ping() error {
	attachment := slack.Attachment{
		Text: "I'm alive!",
	}
	return svc.sendMessage(attachment)
}

func (svc *SlackService) sendMessage(attachment slack.Attachment) error {
	api := slack.New(svc.oauthToken)
	_, _, err := api.PostMessage(
		svc.channelID,
		slack.MsgOptionText(
			"Sifaka spotted:", false),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
	)
	return err
}

func certWarningMessageText(certs []storage.CertificateData) string {
	buf := bytes.NewBufferString("Certificates that will expire soon:\n")
	for _, cert := range certs {
		buf.WriteString(fmt.Sprintf("%s that expires in %s\n", cert.Name, cert.ExpiresInString()))
	}
	return buf.String()
}

func slackProviderEnabled() bool {
	if config.SlackOAuth() != "" && config.SlackChannelID() != "" {
		return true
	}
	return false
}
