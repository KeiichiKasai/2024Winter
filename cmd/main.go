package main

import (
	"2024Winter/app/api/router"
	"2024Winter/initialize"
)

func main() {
	initialize.SetupViper()
	initialize.SetupZap()
	initialize.InitDatabase()
	router.InitRouter()
}
