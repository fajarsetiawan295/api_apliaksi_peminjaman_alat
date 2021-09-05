package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fajars295/api_apliaksi_peminjaman_alat/api/controllers"
	"github.com/fajars295/api_apliaksi_peminjaman_alat/api/services"
	"github.com/joho/godotenv"
)

func main() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("log public path" + pwd)
	fmt.Println("log public dari public" + services.Pathpublic())

	// if err := godotenv.Load(pwd, "/.env"); err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app := controllers.App{}
	app.Initialize(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"))

	app.RunServer()
}
