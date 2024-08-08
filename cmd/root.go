/*
Copyright © 2024 Meiro
*/

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/spf13/cobra"
)

var name, email string

var cliCommand = &cobra.Command{
	Use:   "changegit",
	Short: "Change git's user and/or email",
	Long: `
Important!
If your name has space like: John Wick,
you must pass with the value with "".

	Example:
		changegit -n "John Wick" -e johnwick@bol.com`,

	Run: func(cmd *cobra.Command, args []string) {
		if name == "" && email == "" {
			fmt.Println("Nothing to change.")
			return
		}

		if email != "" {
			emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`)
			if !emailRegex.MatchString(email) {
				fmt.Println("Invalid e-mail, please check again your e-mail input.")
				return
			}

			cmd := exec.Command("git", "config", "--global", "user.email", email)
			err := cmd.Run()
			if err != nil {
				fmt.Println("Someting went wrong.")
			}
		}

		if name != "" {
			cmd := exec.Command("git", "config", "--global", "user.name", name)
			err := cmd.Run()
			if err != nil {
				fmt.Println("Someting went wrong.")
			}
		}
		fmt.Println("Git credentials changed.")
	},
}

func Execute() {
	err := cliCommand.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cliCommand.Flags().StringVarP(&name, "name", "n", "", "user.name value")
	cliCommand.Flags().StringVarP(&email, "email", "e", "", "user.email value")
}
