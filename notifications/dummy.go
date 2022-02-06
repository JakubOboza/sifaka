package notifications

import (
	"fmt"

	"github.com/JakubOboza/sifaka/config"
	"github.com/JakubOboza/sifaka/storage"
)

type DummyService struct {
}

func NewDummyProvider() Provider {
	return &DummyService{}
}

func (svc *DummyService) Notify(certs []storage.CertificateData) error {
	fmt.Println("This certs require refresh / update")
	for _, cert := range certs {
		fmt.Println(cert.ID, cert.Name, cert.Type, cert.ExpiresInString())
	}
	return nil
}

func (svc *DummyService) Ping() error {
	fmt.Println("Dummy notification service pong!")
	return nil
}

func dummyProviderEnabled() bool {
	return config.DummyNotificationsEnabled()
}
