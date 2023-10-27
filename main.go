package main

import (
	"itmx-test/controller"
	"itmx-test/database"
)

func main() {

	database.Init()
	controller.Api()
}
