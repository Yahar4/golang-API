package main

import (
	"database/sql"
	"gihub.com/Yahar4/cmd/api"
	"gihub.com/Yahar4/config"
	"gihub.com/Yahar4/db"
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// DB base connection string with ENV variables
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	// initializing DB inside of main(), so again we can use it
	initStorage(db)

	// server starts
	s := api.NewAPIServer(":9090", db)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}

// here we do implement DB right into our code so we can use it
func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected")
}
