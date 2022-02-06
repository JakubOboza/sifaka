package storage

import (
	"embed"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations/*.sql
var fs embed.FS

func (repo *SqliteRepository) Migrate(databaseURL string) error {
	// Implement database migrations here with embedded  files
	driver, err := iofs.New(fs, "migrations")

	if err != nil {
		return nil
	}

	m, err := migrate.NewWithSourceInstance("iofs", driver, fmt.Sprintf("sqlite3://%s", databaseURL))

	if err != nil {
		return err
	}

	err = m.Up()

	if err != nil {
		return err
	}

	return nil
}
