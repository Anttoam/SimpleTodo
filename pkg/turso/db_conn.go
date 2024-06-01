package turso

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Anttoam/SimpleTodo/config"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func NewLibsqlDB(cfg *config.Config) (*sql.DB, error) {
	url := fmt.Sprintf("libsql://%s.turso.io?authToken=%s", cfg.Turso.Name, cfg.Turso.Token)

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s", err)
		os.Exit(1)
	}
	return db, nil
}
