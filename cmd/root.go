package cmd

import (
	"fmt"
	"os"

	config "github.com/aituglo/go-boilerplate/pkg"
	"github.com/spf13/cobra"
)

var rootCMD = cobra.Command{
	Use:     "go-boilerplate",
	Short:   "go-boilerplate",
	Long:    `go-boilerplate is a boilerplate.`,
	Version: "0.1",
	Run:     startProgram,
}

var conf config.Config

func init() {
	rootCMD.Flags().BoolVarP(&conf.Verbose, "verbose", "v", false, "Verbose")
}

func startProgram(cmd *cobra.Command, args []string) {
	fmt.Println(conf)
}

func Start() {
	if err := rootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
