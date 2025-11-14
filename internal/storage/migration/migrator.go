package migration

type Migrator interface {
	ApplyMigrations() error
}
