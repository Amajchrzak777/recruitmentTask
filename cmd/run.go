package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
)

const FILE_EXTENSION string = ".html"

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "runs http server",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var fileCmd = &cobra.Command{
	Use: "file",
	Short: "add file",
	Run: func(cmd *cobra.Command, args []string) {
		ChangeFileLocation(args[0])
		http.Handle("/", http.FileServer(http.Dir("./static")))
		http.ListenAndServe(":8080", nil)
	},
}

func ChangeFileLocation(arg string) {
	checkPassedArgument(arg)

	oldLocation := arg
	newLocation := "static/" + arg
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
}

func checkPassedArgument(arg string) {
	if arg == "" {
		log.Fatalln("you haven't passed an argument")
	} else if arg[len(arg)-5:] != FILE_EXTENSION {
		log.Fatalln("wrong file extension.")
	} else if len(arg) < 5 {
		log.Fatalln("too short name of file.")
	}
}

func init() {
	runCmd.AddCommand(fileCmd)
	RootCmd.AddCommand(runCmd)
}
