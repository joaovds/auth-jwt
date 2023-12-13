package main

import (
	"fmt"

	"github.com/joaovds/auth-jwt/internals/main/config"
)

func main() {
	config.LoadEnv()
	fmt.Print(config.ENV)
}
