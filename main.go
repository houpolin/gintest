package main

import (
	"gintest/router"
)

func main() {
	router := router.SetupRouter()
	router.Run()
}
