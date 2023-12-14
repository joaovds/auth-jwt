package config

import (
	database "github.com/joaovds/auth-jwt/internal/infra/mongo"
)

func SetupDatabase() {
	database.SetupMongoDB(ENV.DBURI, ENV.DBName)
}
