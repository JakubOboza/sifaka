package watcher

import (
	"testing"

	"github.com/JakubOboza/sifaka/notifications"
	"github.com/JakubOboza/sifaka/storage"
)

func TestWatcher(t *testing.T) {

	notficationService := notifications.NewTestMock(func(certs []storage.CertificateData) error {
		if len(certs) != 3 {
			t.Errorf("Expected 3 certs sent to notifications but got'%v'", len(certs))
		}
		return nil
	})

	storageService := storage.NewTestMock()

	wat := New(storageService, notficationService)

	wat.runSinglePass()

}
