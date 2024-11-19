package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "image-api",
	Short: "obsidian image-api gateway",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpTemplate()
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
