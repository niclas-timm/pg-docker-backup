package main

import (
	"pg-docker-backup/cmd"
	"pg-docker-backup/notifications"

	"github.com/joho/godotenv"
)

// main is the entry method of the program.
func main(){

	err := godotenv.Load()
	if err != nil {
		notifications.NotifyViaAllChannels("Unable to load environment variables. Aborting backup.")
		panic("Could not load environment variables")
	}
	cmd.Execute()
}