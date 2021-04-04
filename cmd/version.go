package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version app",
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}

func printVersion() {
	goVersion := runtime.Version()
	appVersion := "1.0.0"
	fmt.Println("GoLang version:\t", goVersion)
	fmt.Println("App version:\t", appVersion)

}

func init() {
	RootCmd.AddCommand(versionCmd)
}