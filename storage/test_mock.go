package storage

import (
	"time"

	"github.com/JakubOboza/sifaka/certutils"
)

type StorageTestMockService struct {
}

func NewTestMock() Service {
	return &StorageTestMockService{}
}

func (mock *StorageTestMockService) FindById(id int) (*CertificateData, error) {
	return &CertificateData{
		ID:        id,
		Name:      "TestCertData.cer",
		Subject:   "CN:Mock",
		Type:      "file",
		NotAfter:  time.Now().Add(1 * time.Hour),
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}, nil
}

func (mock *StorageTestMockService) AllByExpiration() ([]CertificateData, error) {
	return []CertificateData{
		{
			ID:        420,
			Name:      "TestCertData.cer",
			Subject:   "CN:Mock",
			Type:      certutils.TYPE_FILE,
			NotAfter:  time.Now().Add(1 * time.Hour),
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		}, {
			ID:        69,
			Name:      "https://lambdacu.be",
			Subject:   "CN:Mock",
			Type:      certutils.TYPE_URL,
			NotAfter:  time.Now().Add(1 * time.Hour),
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		},
	}, nil
}

func (mock *StorageTestMockService) AllForIndexPage() ([]DisplayData, error) {
	cd := &CertificateData{
		ID:        69,
		Name:      "TestCertData.cer",
		Subject:   "CN:Mock",
		Type:      "file",
		NotAfter:  time.Now().Add(1 * time.Hour),
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	return []DisplayData{cd.ToDisplayData()}, nil
}

func (mock *StorageTestMockService) Store(createCertData *CreateCertificateData) (*CertificateData, error) {
	return mock.FindById(777)
}

func (mock *StorageTestMockService) Destroy(certData *CertificateData) error {
	return nil
}

func (mock *StorageTestMockService) Update(certData *CertificateData) error {
	return nil
}

func (mock *StorageTestMockService) AllWebCertsThatMightExpireSoon() ([]CertificateData, error) {
	return []CertificateData{
		{
			ID:        420,
			Name:      "https://google.com",
			Subject:   "CN:Mock",
			Type:      certutils.TYPE_URL,
			NotAfter:  time.Now().Add(1 * time.Hour),
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		}, {
			ID:        69,
			Name:      "https://lambdacu.be",
			Subject:   "CN:Mock",
			Type:      certutils.TYPE_URL,
			NotAfter:  time.Now().Add(1 * time.Hour),
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		},
	}, nil
}

func (mock *StorageTestMockService) AllWCertsThatExpireSoon() ([]CertificateData, error) {
	return []CertificateData{
		{
			ID:        420,
			Name:      "https://google.com",
			Subject:   "CN:Mock",
			Type:      certutils.TYPE_URL,
			NotAfter:  time.Now().Add(1 * time.Hour),
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		}, {
			ID:        69,
			Name:      "https://lambdacu.be",
			Subject:   "CN:Mock",
			Type:      certutils.TYPE_URL,
			NotAfter:  time.Now().Add(1 * time.Hour),
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		},
		{
			ID:        1410,
			Name:      "TestCertData.cer",
			Subject:   "CN:Mock",
			Type:      certutils.TYPE_FILE,
			NotAfter:  time.Now().Add(1 * time.Hour),
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		},
	}, nil
}
