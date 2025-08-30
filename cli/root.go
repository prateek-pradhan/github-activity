package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
    Use:   "github-activity",
    Short: "A simple CLI to fetch events from GitHub for a user",
}

var githubActivityCmd = &cobra.Command{
	Use:   "github-activity [username]",
	Short: "Fetch and display GitHub activity",
	Long:  `Fetch and display GitHub activity such as commits, pull requests, and issues.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
            fmt.Println("Please provide a GitHub username")
            return
        }
		getGithubActivityOfUser(args[0])
	},
}

func Execute() {
	rootCmd.AddCommand(githubActivityCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error occured while executing Github Activity'%s'\n", err)
		os.Exit(1)
	}
}