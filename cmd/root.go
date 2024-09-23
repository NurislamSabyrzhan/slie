package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "slie",
	Short: "CLI tool for generating newserver project files",
	Long:  `A CLI tool to generate router, controller, and usecase files for the newserver project.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}