package storage

import (
	"os"
	"testing"
	"time"

	"github.com/JakubOboza/sifaka/certutils"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const (
	TEST_SQLITE_DB_FILE = "../test/test.db"
)

func TestSqliteRepository(t *testing.T) {

	os.Remove(TEST_SQLITE_DB_FILE)

	conn, err := sqlx.Connect("sqlite3", TEST_SQLITE_DB_FILE)
	if err != nil {
		t.Errorf("Unexpected error when bootstraping sqlite for tests: '%v'", err)
		return
	}
	defer conn.Close()

	repo := NewSqliteRepository(conn)

	err = repo.Migrate(TEST_SQLITE_DB_FILE)

	if err != nil {
		t.Errorf("Error while migrating sqlite db for tests: '%v'", err)
	}

	notAfter := time.Now()

	certData, err := repo.Store(&CreateCertificateData{
		Name:     "Test Cert Data",
		Subject:  "Sifaka that cert?",
		Type:     certutils.TYPE_FILE,
		NotAfter: notAfter,
	})

	if err != nil {
		t.Errorf("repo.Store() unexpected error: '%v'", err)
	}

	timeDiff := certData.NotAfter.Sub(notAfter)

	if timeDiff > 0 {
		t.Errorf("time not_after stored is different to not_after fetched by '%v'", timeDiff)
	}

	certData.Subject = "New Subject"
	err = repo.Update(certData)

	if err != nil {
		t.Errorf("repo.Update() unexpected error: '%v'", err)
	}

	certDataReload, err := repo.FindById(certData.ID)

	if err != nil {
		t.Errorf("repo.FindById() unexpected error: '%v'", err)
	}

	if certDataReload.Subject != "New Subject" {
		t.Errorf("repo.Update() didn't update subject")
	}

	_, err = repo.Store(&CreateCertificateData{
		Name:     "https://lambdacu.be",
		Subject:  "Sifaka that cert will expire?",
		Type:     certutils.TYPE_URL,
		NotAfter: notAfter,
	})

	if err != nil {
		t.Errorf("repo.Store() unexpected error: '%v'", err)
	}

	webCerts, err := repo.AllWebCertsThatMightExpireSoon()

	if err != nil {
		t.Errorf("repo.AllWebCertsThatMightExpireSoon() unexpected error: '%v'", err)
	}

	if len(webCerts) != 1 {
		t.Errorf("Expected 1 web cert to expired but got '%v'", len(webCerts))
	}

	willExpireCerts, err := repo.AllWCertsThatExpireSoon()

	if err != nil {
		t.Errorf("repo.AllWCertsThatExpireSoon() unexpected error: '%v'", err)
	}

	if len(willExpireCerts) != 2 {
		t.Errorf("Expected 2 web cert to expired but got '%v'", len(willExpireCerts))
	}

	err = repo.Destroy(certData)

	if err != nil {
		t.Errorf("repo.Destroy() unexpected error: '%v'", err)
	}

	shouldNotExist, err := repo.FindById(certData.ID)

	if shouldNotExist != nil {
		t.Errorf("cert should be removed but wasn't '%v'", shouldNotExist)
	}

	if err.Error() != "sql: no rows in result set" {
		t.Errorf("expected no rows error but got '%v'", err)
	}

}
