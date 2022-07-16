package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

// init method. Used for adding commands to the root command.
func init(){
	rootCmd.AddCommand(backup)
	rootCmd.AddCommand(imp)
}

func Execute() error {
	return rootCmd.Execute()
}