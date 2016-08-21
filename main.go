package main

import (
	"os"

	"github.com/dnem/ci-web/server"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	s := server.NewServer()
	s.Run(":" + port)
}
