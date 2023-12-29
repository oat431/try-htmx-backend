package main

import (
	"github.com/oat431/try-htmx-backend/routes"
)

func main() {
	server := routes.SetupRouter()
	server.Listen(":3000")
}
