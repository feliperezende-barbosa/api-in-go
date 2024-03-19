package database

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type DbMySqlHandler struct {
	db *sql.DB
}

func (dmsh *DbMySqlHandler) Conn() {
	mysqlCfg := mySqlConfig()
	database, err := sql.Open("mysql", mysqlCfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	dmsh.db = database
}

func mySqlConfig() *mysql.Config {
	cfg := mysql.Config{
		User:   "",
		Passwd: "",
		Net:    "tcp",
		Addr:   "",
		DBName: "albums",
	}
	return &cfg
}
