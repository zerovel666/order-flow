package migration

import (
	"context"
	"database/sql"
	"embed"
	"log"

	"github.com/pressly/goose/v3"
)

//go:embed *.sql
var embedFS embed.FS

func RunMigrations(ctx context.Context, db *sql.DB) error {
	p, err := goose.NewProvider(goose.DialectPostgres, db, embedFS)
	if err != nil {
		return err
	}

	results, err := p.Up(ctx)
	if err != nil {
		return err
	}
	for _, r := range results {
		log.Printf("migration applied: %s (%s)", r.Source.Path, r.Duration)
	}
	return nil
}
