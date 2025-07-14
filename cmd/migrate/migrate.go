package main

import (
	"errors"
	"os"

	"github.com/SaTeR151/EduMinin/internal/config"
	"github.com/SaTeR151/EduMinin/internal/database/sqlite"
	"github.com/golang-migrate/migrate"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Error(err)
		return
	}
	sqliteConfig := config.GetSQLiteConfig()
	db, close, err := sqlite.Open(sqliteConfig)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer close()
	if len(os.Args) < 2 {
		logrus.Fatal("Usage: migrate <command>\nAvailable commands: up, down")
	}
	command := os.Args[1]
	switch command {
	case "up":
		if err := db.MigrationUP(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			logrus.Fatalf("Failed to apply migrations: %v", err)
		}
		logrus.Printf("Migrations applied successfully!")
	case "down":
		if err := db.MigrationDown(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			logrus.Fatalf("Failed to rollback back migrations: %v", err)
		}
		logrus.Printf("Migrations rolled back successfully!")

	default:
		logrus.Fatalf("Unknown command: %s\nAvailable commands: up, down", command)
	}
}
