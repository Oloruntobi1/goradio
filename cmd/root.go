package cmd

import (
	"fmt"
	"github.com/charmbracelet/log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goradio",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello Goradio!")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func handleErr(err error) {
	if err == nil {
		return
	}

	log.Error(err)
	_, _ = fmt.Fprintf(
		os.Stderr,
		"%s %s\n",
		strings.Trim(err.Error(), " \n"),
	)
	os.Exit(1)
}
