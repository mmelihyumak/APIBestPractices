// internal/database/database.go
package database

import (
    "database/sql"
    _ "github.com/lib/pq"
    "time"
	"log"
)

var DB *sql.DB

func InitDB(dataSourceName string) (*sql.DB, error) {
    var err error
    DB, err = sql.Open("postgres", dataSourceName)
    if err != nil {
        return nil, err
    }

    DB.SetMaxOpenConns(25)
    DB.SetMaxIdleConns(25)
    DB.SetConnMaxLifetime(5 * time.Minute)

    if err = DB.Ping(); err != nil {
        DB.Close()
        return nil, err
    }

	log.Println("DB bağlantısı başarılı")
    return DB, nil
}
