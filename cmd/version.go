package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of GoUnderscore",
	Long:  `All software has versions. This is GoUnderscore's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GoUnderscore v0.01 -- HEAD")
	},
}
