package main

import (
	"fmt"

	"github.com/SethHafferkamp/GoRSS/models"
)

func main() {
	fmt.Println("hello, world\n")
	fmt.Println("Dropping all models")
	models.MigrateToZero()
	fmt.Println("All Done!!")
}
