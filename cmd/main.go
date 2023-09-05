package main

import (
	"fmt"
	"log"
	"numfingpt/pkg"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// example of usage
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	epam_url := os.Getenv("EPAM_URL")
	amzn_url := os.Getenv("AMZN_URL")
	API_KEY := os.Getenv("GPT_API_KEY")
	pkg.Exec(epam_url, API_KEY)
	fmt.Print("\n\n\n")
	pkg.Exec(amzn_url, API_KEY)
}
