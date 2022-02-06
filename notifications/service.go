package notifications

import (
	"fmt"

	"github.com/JakubOboza/sifaka/storage"
)

type Provider interface {
	Notify(certs []storage.CertificateData) error
	Ping() error
}

type Service interface {
	Notify(certs []storage.CertificateData) error
	Add(provider Provider)
}

type NotificationService struct {
	notificationProviders []Provider
}

func New() Service {
	notificationService := &NotificationService{
		notificationProviders: []Provider{},
	}

	if dummyProviderEnabled() {
		notificationService.Add(NewDummyProvider())
	}

	if slackProviderEnabled() {
		slackService := NewSlackProvider()
		err := slackService.Ping()
		if err == nil {
			fmt.Println("Slack service ping err:", err)
		} else {
			notificationService.Add(slackService)
		}
	}

	return notificationService
}

func (svc *NotificationService) Notify(certs []storage.CertificateData) error {
	for _, service := range svc.notificationProviders {
		service.Notify(certs)
	}
	return nil
}

func (svc *NotificationService) Add(provider Provider) {
	svc.notificationProviders = append(svc.notificationProviders, provider)
}
