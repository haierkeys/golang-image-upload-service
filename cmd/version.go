package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/haierspi/golang-image-upload-service/global"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print out version info and exit.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("v%s ( Git:%s ) BuidTime:%s\n", global.Version, global.GitTag, global.BuildTime)

	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
