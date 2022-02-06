package notifications

import "github.com/JakubOboza/sifaka/storage"

type NotificationTestMockService struct {
	notifyHandler func(certs []storage.CertificateData) error
}

func NewTestMock(notifyHandler func(certs []storage.CertificateData) error) Service {
	return &NotificationTestMockService{notifyHandler: notifyHandler}
}

func (mock *NotificationTestMockService) Notify(certs []storage.CertificateData) error {
	return mock.notifyHandler(certs)
}

func (mock *NotificationTestMockService) Add(provider Provider) {

}
