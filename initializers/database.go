package initializers

import (
    "os"
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToSqliteDB() {
    var err error
    dsn := os.Getenv("SQLITE_DB_URL")
    DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database")
    }
}

func ConnectToPostgresDB() {
    var err error
    dsn := os.Getenv("POSTGRES_DB_URL")
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database")
    }
}
