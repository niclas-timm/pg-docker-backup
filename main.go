package main

import (
	"github.com/joho/godotenv"

	"github.com/NiclasTimmeDev/pg-docker-backup/cmd"
	"github.com/NiclasTimmeDev/pg-docker-backup/notifications"
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