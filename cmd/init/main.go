package main

import (
	"fmt"

	"github.com/SethHafferkamp/GoRSS/helpers"
	"github.com/SethHafferkamp/GoRSS/models"
)

func main() {
	fmt.Println("hello, world\n")
	fmt.Println("Running migrations")
	models.RunMigrations()
	fmt.Println("Register Test Data. This will take a couple minutes")
	helpers.RegisterTestData()
	fmt.Println("All Done!!")
}
