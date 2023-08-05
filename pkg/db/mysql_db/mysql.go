package mysql_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConfig struct {
	Username     string `yaml:"user"`
	Password     string `yaml:"password"`
	Port         string `yaml:"port"`
	DatabaseName string `yaml:"database"`
	Host         string `yaml:"host"`
}

func ConnectDatabase(args DatabaseConfig) (*sql.DB, error) {
	var con = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", args.Username, args.Password, args.Host, args.Port, args.DatabaseName)
	db, err := sql.Open("mysql", con)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	return db, nil
}
