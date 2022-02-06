/*
Copyright © 2022 Jakub Oboza <jakub.oboza@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package certutils

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"net/url"
	"time"
)

const (
	TYPE_URL  = "url"
	TYPE_FILE = "file"
)

var (
	ErrCertNotFound                 = errors.New("Certificate for a given url couldn't be found or matched")
	ErrMoreThanOneNoneCaCertInChain = errors.New("Certificate chain has more than one non CA cert. That is a problem")
)

type CertInfo struct {
	Type        string    `json:"type"`
	Subject     string    `json:"subject"`
	NotAfter    time.Time `json:"not_after"`
	Certificate *x509.Certificate
}

func (ci *CertInfo) ExpiresIn() time.Duration {
	return ci.NotAfter.Sub(time.Now())
}

func (ci *CertInfo) ExpiresInString() string {
	days := ci.ExpiresIn().Hours() / 24
	if days < 1.0 && days > -1.0 {
		return fmt.Sprintf("%s", ci.ExpiresIn())
	}
	return fmt.Sprintf("%.0f days", days)
}

func (ci *CertInfo) Status() string {
	if ci.ExpiresIn() < 0 {
		return "expired"
	}
	return "active"
}

func (ci *CertInfo) DisplayInfo() {
	fmt.Println("====CERT INFO====")
	fmt.Println("Type:", ci.Type)
	fmt.Println("Status:", ci.Status())
	fmt.Println("Subject:", ci.Subject)
	fmt.Println("Expires at", ci.NotAfter)
	fmt.Println("Expires in", ci.ExpiresInString())
	fmt.Println("")
}

func GetCertInfoForFile(rawCertBlockFromFile []byte) (*CertInfo, error) {
	block, _ := pem.Decode(rawCertBlockFromFile)

	cert, err := x509.ParseCertificate(block.Bytes)

	if err != nil {
		return nil, err
	}

	return &CertInfo{
		NotAfter:    cert.NotAfter,
		Subject:     cert.Subject.String(),
		Certificate: cert,
		Type:        TYPE_FILE,
	}, nil

}

func GetCertInfoForUrl(urlInfo *url.URL) (*CertInfo, error) {
	port := "443"
	//For port 80 we check 443 for cert but
	//for set fixed ports we use provided data
	if urlInfo.Port() != "" {
		port = urlInfo.Port()
	}

	addr := fmt.Sprintf("%s:%s", urlInfo.Hostname(), port)
	conn, err := tls.Dial("tcp", addr, &tls.Config{
		InsecureSkipVerify: true,
	})

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cert, err := findCertForDomain(urlInfo, conn.ConnectionState().PeerCertificates)

	if err != nil {
		return nil, err
	}

	return &CertInfo{
		NotAfter:    cert.NotAfter,
		Subject:     cert.Subject.String(),
		Certificate: cert,
		Type:        TYPE_URL,
	}, nil
}

func nonCaCertsCount(certs []*x509.Certificate) (count int, index int) {
	for idx, cert := range certs {
		if !cert.IsCA {
			count += 1
			index = idx
		}
	}
	return
}

func findCertForDomain(urlPath *url.URL, certs []*x509.Certificate) (*x509.Certificate, error) {
	countOfNonCaCerts, index := nonCaCertsCount(certs)
	if countOfNonCaCerts == 1 {
		return certs[index], nil
	} else if countOfNonCaCerts > 1 {
		//TODO: do research on this case
		return nil, ErrMoreThanOneNoneCaCertInChain
	} else {
		return nil, ErrCertNotFound
	}
}
