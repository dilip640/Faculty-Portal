package storage

import (
	migrate "github.com/rubenv/sql-migrate"
	log "github.com/sirupsen/logrus"
)

// Migrate migrate the database all the way up
func Migrate() {
	Initialize()

	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Applied %d migrations!\n", n)

}

//DownOneStep will migrate one version down
func DownOneStep() {
	Initialize()

	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	n, err := migrate.ExecMax(db, "postgres", migrations, migrate.Down, 1)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Applied %d migrations!\n", n)
}
