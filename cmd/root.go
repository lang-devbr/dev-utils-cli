package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dev-utils-cli",
	Short: "dev utils cli root command",
	Long:  "dev utils cli root command",
}

// Execute root command
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal("error while executing command")
	}
}
