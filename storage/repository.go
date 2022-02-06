package storage

type Migrator interface {
	Migrate(databaseURL string) error
}

type Reader interface {
	FindById(id int) (*CertificateData, error)
	AllByExpiration() ([]CertificateData, error)
	AllWebCertsThatMightExpireSoon() ([]CertificateData, error)
	AllWCertsThatExpireSoon() ([]CertificateData, error)
}

type Writer interface {
	Store(createCertData *CreateCertificateData) (*CertificateData, error)
	Destroy(certData *CertificateData) error
	Update(certData *CertificateData) error
}

type Repository interface {
	Migrator
	Reader
	Writer
}
