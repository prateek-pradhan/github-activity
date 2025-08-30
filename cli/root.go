package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var githubActivityCmd = &cobra.Command{
	Use:   "github-activity",
	Short: "Fetch and display GitHub activity",
	Long:  `Fetch and display GitHub activity such as commits, pull requests, and issues.`,
	Run: func(cmd *cobra.Command, args []string) {
		getGithubActivityOfUser(args[1])
	},
}

func Execute() {
	if err := githubActivityCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Oops. An error occured while executing Github Activity'%s'\n", err)
		os.Exit(1)
	}
}