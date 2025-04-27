package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "0.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of change git",
	Long:  `All software has versions. This is change git's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("changegit v%s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
