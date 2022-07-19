package main

import (
	"airi_back/models"
	. "airi_back/router"
)

func main() {
	models.DB()
	Router()
}
