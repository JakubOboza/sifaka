package storage

import (
	"fmt"
	"time"

	"github.com/JakubOboza/sifaka/certutils"
)

type CertificateData struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Subject   string    `db:"subject" json:"subject"`
	Type      string    `db:"type" json:"type"`
	NotAfter  time.Time `db:"not_after" json:"not_after"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

func New(name string, certInfo *certutils.CertInfo) *CertificateData {
	return &CertificateData{
		Name:     name,
		Subject:  certInfo.Subject,
		Type:     certInfo.Type,
		NotAfter: certInfo.NotAfter,
	}
}

func (cd *CertificateData) ExpiresIn() time.Duration {
	return cd.NotAfter.Sub(time.Now())
}

func (cd *CertificateData) ExpiresInString() string {
	days := cd.ExpiresIn().Hours() / 24
	if days < 1.0 && days > -1.0 {
		return fmt.Sprintf("%s", cd.ExpiresIn())
	}
	return fmt.Sprintf("%.0f days", days)
}

func (cd *CertificateData) Status() string {
	if cd.ExpiresIn() < 0 {
		return "expired"
	}
	return "active"
}

//CreateCertificateData used to create CertificateData record
type CreateCertificateData struct {
	Name     string
	Subject  string
	Type     string
	NotAfter time.Time
}

//DisplayData has all data in final concrete format ready for display in eg. HTML template
type DisplayData struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Subject   string    `json:"subject"`
	Type      string    `json:"type"`
	ExpiresIn string    `json:"expires_in"`
	NotAfter  time.Time `json:"not_after"`
	Status    string    `json:"status"`
	Button    string    `json:"status"`
}

func (cd *CertificateData) ToDisplayData() DisplayData {
	return DisplayData{
		ID:        cd.ID,
		Name:      cd.Name,
		Subject:   cd.Subject,
		Type:      cd.Type,
		NotAfter:  cd.NotAfter,
		ExpiresIn: cd.ExpiresInString(),
		Status:    cd.Status(),
		Button:    buttonColor(cd.NotAfter),
	}
}

// Less than 30 days warning
// Expired = Danger
func buttonColor(t time.Time) string {
	timeDiff := t.Sub(time.Now())
	if timeDiff < 0 {
		return "danger"
	} else if timeDiff < 30*24*time.Hour {
		return "warning"
	}
	return "success"
}
