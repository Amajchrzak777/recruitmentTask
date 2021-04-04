package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command {
	Use: "recruitment task",
	Short: "recruitment task, CLI app which starts http server with html file provided by user",
}