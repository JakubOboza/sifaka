-- type CertificateData struct {
-- 	ID        int       `db:"id" json:"id"`
-- 	Name      string    `db:"name" json:"name"`
-- 	Subject   string    `db:"subject" json:"subject"`
-- 	Type      string    `db:"type" json:"type"`
-- 	NotAfter  time.Time `db:"not_after" json:"not_after"`
-- 	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
-- 	CreatedAt time.Time `db:"created_at" json:"created_at"`
-- }

--Driver automatically wraps every migration in transaction
--BEGIN;

CREATE TABLE IF NOT EXISTS certs(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT,
  subject TEXT,
  type TEXT,
  not_after  DATETIME,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

--COMMIT;