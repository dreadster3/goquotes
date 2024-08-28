package database

import (
	"database/sql"
	"errors"
	"os"

	"github.com/dreadster3/goquotes/database/generated"
	_ "github.com/mattn/go-sqlite3"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	ErrDriverNotSpecified     = errors.New("database driver not specified")
	ErrDataSourceNotSpecified = errors.New("database data source not specified")
)

var (
	Queries  *generated.Queries
	Database *sql.DB
)

func InitDB() error {
	driver := os.Getenv("DB_DRIVER")
	dataSource := os.Getenv("DB_DATA_SOURCE")
	if driver == "" && dataSource == "" {
		driver = "sqlite3"
		dataSource = "./test.db"
	}

	if driver == "" {
		return ErrDriverNotSpecified
	}

	if dataSource == "" {
		return ErrDataSourceNotSpecified
	}

	db, err := sql.Open(driver, dataSource)
	if err != nil {
		return err
	}

	dbDriver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://database/migrations/", "goquotes", dbDriver)
	if err != nil {
		return err
	}

	if _, _, err = m.Version(); err != nil && err == migrate.ErrNilVersion {
		err = m.Up()
		if err != nil {
			return err
		}
	}

	Database = db
	Queries = generated.New(Database)

	return nil
}
