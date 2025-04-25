package main

import (
	"backend/rewards-recognition-api/database"
	"backend/rewards-recognition-api/handler"
	"fmt"
	"log"

	"github.com/common-nighthawk/go-figure"
	"rsc.io/quote"
)

func main() {
	log.Println("Starting Application")
	myFigure := figure.NewColorFigure("Rewards and Recognition", "", "green", true)
	myFigure.Print()
	fmt.Println(quote.Go())

	db := &database.Database{}
	routers := handler.Routers{}

	db.Initialize()
	routers.Initialize(":8080", db.DB)
}
