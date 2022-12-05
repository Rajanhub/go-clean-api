package infrastructure

import (
	"github.com/Rajanhub/goapi/lib"
	migrate "github.com/rubenv/sql-migrate"
)

type Migration interface {
	Migrate() error
}

type migration struct {
	logger lib.Logger
	db     Database
}

func NewMigration(loggger lib.Logger, db Database) Migration {
	return &migration{
		logger: loggger,
		db:     db,
	}
}

func (m *migration) Migrate() error {
	migration := &migrate.FileMigrationSource{Dir: "migration/"}

	sqlDB, err := m.db.DB.DB()
	if err != nil {
		return err
	}

	m.logger.Info("running migration.")
	_, err = migrate.Exec(sqlDB, "mysql", migration, migrate.Up)
	if err != nil {
		return err
	}
	m.logger.Info("migration completed.")
	return nil
}

// RunMigration runs the migration provided logger and database instance
func RunMigration(logger lib.Logger, db Database) error {
	m := &migration{logger, db}
	return m.Migrate()
}
