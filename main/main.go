package main

import (
	"backend/rewards-recognition-api/database"
	"backend/rewards-recognition-api/handler"
	"fmt"
	"log"

	"rsc.io/quote"
)

func main() {
	log.Println("Starting Application")
	log.Println()
	fmt.Println(quote.Go())

	db := &database.Database{}
	routers := handler.Routers{}

	db.Initialize()
	routers.Initialize(":8080", db.DB)
}
