package cmd

import (
	"fmt"
	"github.com/charmbracelet/log"
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "goradio",
	Short: "A brief description of your application",
	Long: `
                               ___     
   ____ _____  _________ _____/ (_)___ 
  / __  / __ \/ ___/ __  / __  / / __ \
 / /_/ / /_/ / /  / /_/ / /_/ / / /_/ /
 \__, /\____/_/   \__,_/\__,_/_/\____/ 
/____/
`,
}

func Execute() {
	cc.Init(&cc.Config{
		RootCmd:         rootCmd,
		Headings:        cc.HiBlue + cc.Bold + cc.Underline,
		Commands:        cc.HiYellow + cc.Bold,
		Example:         cc.Italic,
		ExecName:        cc.Bold,
		Flags:           cc.Bold,
		FlagsDataType:   cc.Italic + cc.HiBlue,
		NoExtraNewlines: true,
		NoBottomNewline: true,
	})

	_ = rootCmd.Execute()
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
		"%s\n",
		strings.Trim(err.Error(), " \n"),
	)
	os.Exit(1)
}
