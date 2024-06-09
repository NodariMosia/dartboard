package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tursodatabase/go-libsql"
)

func main() {
	dbName := "local.db"
	primaryUrl := "libsql://[DATABASE].turso.io"
	authToken := "..."

	dir, err := os.MkdirTemp("", "libsql-embedded")
	if err != nil {
		fmt.Println("Error creating temporary directory:", err)
		os.Exit(1)
	}
	defer os.RemoveAll(dir)

	dbPath := filepath.Join(dir, dbName)

	connector, err := libsql.NewEmbeddedReplicaConnector(dbPath, primaryUrl,
		libsql.WithAuthToken(authToken),
	)
	if err != nil {
		fmt.Println("Error creating connector:", err)
		os.Exit(1)
	}
	defer connector.Close()

	db := sql.OpenDB(connector)
	defer db.Close()
}
