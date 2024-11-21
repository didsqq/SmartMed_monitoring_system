package main

import (
	"os"

	"github.com/didsqq/todo-app/pkg/repository"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {

	}
}
