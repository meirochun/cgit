package cmd

import "github.com/spf13/cobra"

var getCredentialsCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(getCredentialsCmd)
}
