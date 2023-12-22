package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Port   string
	DBName string
	DBURI  string
}

var ENV *Env

func newEnv() *Env {
	port := "3333"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

  return &Env{
		Port:   port,
		DBName: os.Getenv("DBNAME"),
		DBURI:  os.Getenv("DBURI"),
	}
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	ENV = newEnv()
}
