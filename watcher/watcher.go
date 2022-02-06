package watcher

import (
	"fmt"
	"net/url"
	"time"

	"github.com/JakubOboza/sifaka/certutils"
	"github.com/JakubOboza/sifaka/notifications"
	"github.com/JakubOboza/sifaka/storage"
)

var (
	WAIT_TIME = 4 * time.Hour
)

type Watcher struct {
	storageService      storage.Service
	notificationService notifications.Service
}

func New(storageService storage.Service, notificationService notifications.Service) *Watcher {
	return &Watcher{storageService: storageService, notificationService: notificationService}
}

func (watcher *Watcher) Start() {
	for {
		//
		watcher.runSinglePass()
		// wait for some time
		time.Sleep(WAIT_TIME)
	}
}

func (watcher *Watcher) tryToUpdateAllCertsForWebsites() error {
	// Websites might need a refresh
	websiteCertsThatMightExpireSoon, err := watcher.storageService.AllWebCertsThatMightExpireSoon()

	if err != nil {
		return err
	}

	for _, certData := range websiteCertsThatMightExpireSoon {
		parsedUrl, err := url.Parse(certData.Name)
		if err != nil {
			continue
		}

		webCert, err := certutils.GetCertInfoForUrl(parsedUrl)

		if err != nil {
			continue
		}

		if webCert.NotAfter.Sub(certData.NotAfter) > 0 {
			// update
			certData.NotAfter = webCert.NotAfter
			certData.Subject = webCert.Subject
			watcher.storageService.Update(&certData)
		}

	}

	return nil
}

func (watcher *Watcher) runSinglePass() {
	err := watcher.tryToUpdateAllCertsForWebsites()

	if err != nil {
		fmt.Println("watcher err:", err)
	}

	certs, err := watcher.storageService.AllWCertsThatExpireSoon()

	if err != nil {
		fmt.Println("watcher err:", err)
	}

	if len(certs) > 0 {
		watcher.notificationService.Notify(certs)
	}
}
