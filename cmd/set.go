package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/spf13/cobra"
)

var (
	email    string
	username string
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set Git user credentials",
	Long:  `Set the global Git username and email credentials`,
	Run: func(cmd *cobra.Command, args []string) {
		if username == "" && email == "" {
			fmt.Println("Please provide at least one of --username or --email")
			os.Exit(1)
		}

		if email != "" {
			if isValidEmail(email) {
				setGitConfig("user.email", email)
			}
		}

		if username != "" {
			setGitConfig("user.name", username)
		}

		fmt.Println("Git credentials updated successfully")
	},
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`)
	if !emailRegex.MatchString(email) {
		fmt.Println("Invalid e-mail, please check again your e-mail input.")
		os.Exit(1)
	}
	return true
}

func setGitConfig(key, value string) {
	cmd := exec.Command("git", "config", "--global", key, value)
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error settings %s: %v\n", key, err)
	}
}

func init() {
	rootCmd.AddCommand(setCmd)

	setCmd.Flags().StringVarP(&username, "username", "u", "", "Git username to set")
	setCmd.Flags().StringVarP(&email, "email", "e", "", "Git email to set")
}
