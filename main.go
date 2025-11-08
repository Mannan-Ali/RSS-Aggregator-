package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hey first a project in GO")

	godotenv.Load() //to load the env file we use this function

	//to use env go has the follwing functoin
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the env")
	}
	fmt.Println("PORT:", portString)
}
