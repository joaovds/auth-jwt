package main

import (
  "fmt"
  "log"

  "github.com/joaovds/auth-jwt/internal/main/config"
)

func main() {
  config.LoadEnv()

  app := config.GetApp()
  fmt.Println("Server running on port", config.ENV.Port)
  log.Fatal(app.Listen(":" + config.ENV.Port))
}
