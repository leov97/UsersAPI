package connection

import (
	"UserAPI/internal/config/data"
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Datasql() *sql.DB {
	credential := data.NewDatabaseConfig()
	cfg := mysql.Config{
		User:   credential.Database.User,
		Passwd: credential.Database.Password,
		Net:    credential.Database.Network,
		Addr:   credential.Database.Address,
		DBName: credential.Database.DBName,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	return db
}
