package storage

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type SqliteRepository struct {
	conn *sqlx.DB
}

func NewSqliteRepository(connectionPool *sqlx.DB) *SqliteRepository {
	return &SqliteRepository{conn: connectionPool}
}

func (repo *SqliteRepository) FindById(id int) (*CertificateData, error) {
	query := "SELECT * FROM certs WHERE id=?"

	certData := &CertificateData{}

	err := repo.conn.Get(certData, query, id)

	if err != nil {
		return nil, err
	}

	return certData, nil
}

func (repo *SqliteRepository) AllByExpiration() ([]CertificateData, error) {
	query := "SELECT * FROM certs ORDER BY not_after ASC"
	certs := []CertificateData{}

	err := repo.conn.Select(&certs, query)

	if err != nil {
		return nil, err
	}

	return certs, nil
}

func (repo *SqliteRepository) Store(certData *CreateCertificateData) (*CertificateData, error) {
	query := "INSERT INTO certs (name, subject, type, not_after, created_at, updated_at) VALUES (?,?,?,?,?,?) RETURNING id"

	tx, err := repo.conn.Begin()

	if err != nil {
		return nil, err
	}

	ID := 0

	err = tx.QueryRow(query, certData.Name, certData.Subject, certData.Type, certData.NotAfter, time.Now(), time.Now()).Scan(&ID)

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return repo.FindById(ID)
}

func (repo *SqliteRepository) Destroy(certData *CertificateData) error {
	query := "DELETE FROM certs WHERE id=?"

	if certData.ID == 0 {
		return ErrRecordIdOutOfBands
	}

	tx, err := repo.conn.Begin()

	if err != nil {
		return err
	}

	tx.Exec(query, certData.ID)

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

func (repo *SqliteRepository) Update(certData *CertificateData) error {
	query := "UPDATE certs SET subject=?, not_after=?, updated_at=? WHERE id=?"

	tx, err := repo.conn.Begin()

	if err != nil {
		return err
	}

	tx.Exec(query, certData.Subject, certData.NotAfter, time.Now(), certData.ID)

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (repo *SqliteRepository) AllWebCertsThatMightExpireSoon() ([]CertificateData, error) {
	cutOfDate := expireDate()
	query := "SELECT * FROM certs WHERE type = \"url\" AND not_after <= ? ORDER BY not_after ASC"
	certs := []CertificateData{}

	err := repo.conn.Select(&certs, query, cutOfDate)

	if err != nil {
		return nil, err
	}

	return certs, nil
}

func (repo *SqliteRepository) AllWCertsThatExpireSoon() ([]CertificateData, error) {
	cutOfDate := expireDate()
	query := "SELECT * FROM certs WHERE  not_after <= ? ORDER BY not_after ASC"
	certs := []CertificateData{}

	err := repo.conn.Select(&certs, query, cutOfDate)

	if err != nil {
		return nil, err
	}

	return certs, nil
}
