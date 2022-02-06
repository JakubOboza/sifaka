package storage

type Service interface {
	FindById(id int) (*CertificateData, error)
	AllByExpiration() ([]CertificateData, error)
	AllForIndexPage() ([]DisplayData, error)
	Store(createCertData *CreateCertificateData) (*CertificateData, error)
	Destroy(certData *CertificateData) error
	Update(certData *CertificateData) error
	AllWebCertsThatMightExpireSoon() ([]CertificateData, error)
	AllWCertsThatExpireSoon() ([]CertificateData, error)
}

type CertDataService struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &CertDataService{repo: repo}
}

func (svc *CertDataService) FindById(id int) (*CertificateData, error) {
	return svc.repo.FindById(id)
}

func (svc *CertDataService) AllByExpiration() ([]CertificateData, error) {
	return svc.repo.AllByExpiration()
}

func (svc *CertDataService) Store(certData *CreateCertificateData) (*CertificateData, error) {
	return svc.repo.Store(certData)
}

func (svc *CertDataService) Destroy(certData *CertificateData) error {
	return svc.repo.Destroy(certData)
}

func (svc *CertDataService) Update(certData *CertificateData) error {
	return svc.repo.Update(certData)
}

func (svc *CertDataService) AllForIndexPage() ([]DisplayData, error) {
	displayableCertInfo := []DisplayData{}

	certsData, err := svc.repo.AllByExpiration()
	if err != nil {
		return nil, err
	}

	for _, certDatum := range certsData {
		displayableCertInfo = append(displayableCertInfo, certDatum.ToDisplayData())
	}

	return displayableCertInfo, nil
}

func (svc *CertDataService) AllWebCertsThatMightExpireSoon() ([]CertificateData, error) {
	return svc.repo.AllWebCertsThatMightExpireSoon()
}

func (svc *CertDataService) AllWCertsThatExpireSoon() ([]CertificateData, error) {
	return svc.repo.AllWCertsThatExpireSoon()
}
