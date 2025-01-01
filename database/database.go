package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	"github.com/pressly/goose/v3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type Config struct {
	TURSO_DATABASE_url string
	TURSO_AUTH_TOKEN   string
}

// LoadEnv reads a JSON file and parses it into the provided config structure.
func LoadEnv(filepath string, config *Config) error {
	projectName := regexp.MustCompile(`^(.*` + "v2" + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	filePath := string(rootPath) + filepath
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open env file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read env file: %w", err)
	}

	err = json.Unmarshal(data, config)
	if err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	return nil
}

func Migrate(db *sql.DB, path string) {
	goose.SetDialect("turso")

	// Apply all "up" migrations
	err := goose.Up(db, path)
	if err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	log.Println("Migrations applied successfully!")
}

func DropMigrate(db *sql.DB, path string) {
	goose.SetDialect("turso")

	// Apply all "up" migrations
	err := goose.Down(db, path)
	if err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	log.Println("Migrations applied successfully!")
}

func ConnectToDB(url, token string) *sql.DB {
	// url := "libsql://[DATABASE].turso.io?authToken=[TOKEN]"
	dburl := fmt.Sprintf("%s?authToken=%s", url, token)

	db, err := sql.Open("libsql", dburl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}

	return db
}
