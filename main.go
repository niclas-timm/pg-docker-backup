package main

import (
	"github.com/NiclasTimmeDev/pg-docker-backup/cmd"
	"github.com/NiclasTimmeDev/pg-docker-backup/config"
	"github.com/NiclasTimmeDev/pg-docker-backup/notifications"
	"github.com/joho/godotenv"
)

// main is the entry method of the program.


func main(){

	err := godotenv.Load()

	config.BootstrapConfig()
	
	if err != nil {
		notifications.NotifyViaAllChannels("Unable to load environment variables. Aborting backup.")
		panic("Could not load environment variables")
	}
	cmd.Execute()
}