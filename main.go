package main

import (
	"flag"
	"fmt"
	"os"
	"promotions-app/config"
	"promotions-app/db"
	"promotions-app/server"
)

func main() {
	environment := getEnvironment()

	config.Init(*environment)
	db.Init()

	server.Start()
}

func getEnvironment() *string {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Flags usage: -e {environment}")
		os.Exit(1)
	}
	flag.Parse()

	return environment
}
