package db

import (
	"context"
	"database/sql"
	"g2r-api/env"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // required
	_ "github.com/lib/pq"                                //Required for postgres driver
	"github.com/myles-mcdonnell/blondie"
	"github.com/myles-mcdonnell/logrusx"
	"github.com/pkg/errors"
	"time"
)

type (
	Database struct {
		*Queries
		db *sql.DB
	}
)

func NewDatabase(db *sql.DB) *Database {
	return &Database{Queries: New(db), db: db}
}

func (database *Database) Ping() error {
	return database.db.Ping()
}

func (database *Database) WithDefaultTransaction(ctx context.Context, fn func(*sql.Tx, *Queries) error) error {

	tx, err := database.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return errors.Wrap(err, "unable to begin tx")
	}

	defer func() {
		if p := recover(); p != nil {
			if err := tx.Rollback(); err != nil {
				logrusx.Error("tx rollback failure").Write(err)
			}
		} else if err != nil {
			logrusx.Error("tx rollback failure").Write(err)
		} else {
			err = tx.Commit()
		}
	}()

	return fn(tx, database.WithTx(tx))
}

func BootstrapDB(config env.PostgresConfig) (*sql.DB, error) {
	blondieOpts := blondie.DefaultOptions()
	blondieOpts.QuietMode = true

	blondie.WaitForDeps([]blondie.DepCheck{
		blondie.NewTcpCheck(config.Host, config.Port, 15*time.Second),
	}, blondieOpts)

	db, err := connect(3, config)
	if err != nil {
		return nil, err
	}

	if config.ApplySchemaMigration {
		db.Close()
		if err := applySchemaMigration(&config); err != nil {
			logrusx.Error("Error applying schema migration").Write(err)
			return nil, err
		}

		db, err = connect(1, config)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

func connect(retries int, config env.PostgresConfig) (*sql.DB, error) {
	var err error
	var db *sql.DB
	for retries > 0 {
		db, err = sql.Open("postgres", config.Address())
		if err == nil {
			err = db.Ping()
			if err == nil {
				break
			}
		}

		logrusx.Infof("Database connection failure; retries left %v", retries).Write(err)

		time.Sleep(1 * time.Second)
		retries--
	}

	if err != nil {
		logrusx.Error("Database connection failure").Write(err)
		return nil, err
	}

	return db, nil
}

// ApplySchemaMigration Bootstrap creates and migrates db schema versions
func applySchemaMigration(config *env.PostgresConfig) error {
	m, err := migrate.New(
		config.MigrationPath,
		config.Address(),
	)
	if err != nil {
		return err
	}
	defer m.Close()

	if config.ResetDB {
		if err := m.Force(0); err != nil {
			return err
		}
		if err = m.Down(); err != nil {
			return err
		}
	}

	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			return err
		}
	}
	_, _, err = m.Version()
	return err
}
