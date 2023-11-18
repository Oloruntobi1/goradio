package cmd

import (
	"github.com/neptunsk1y/goradio/version"
	"html/template"
	"runtime"

	"github.com/charmbracelet/lipgloss"

	"github.com/samber/lo"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)

	versionCmd.Flags().BoolP("short", "s", false, "print the version number only")
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of the goradio",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := version.Latest()
		if err != nil {
			return
		}
		if lo.Must(cmd.Flags().GetBool("short")) {
			_, err := cmd.OutOrStdout().Write([]byte("v" + version.Version + "\n"))
			handleErr(err)
			return
		}

		versionInfo := struct {
			Version  string
			OS       string
			Arch     string
			App      string
			Compiler string
		}{
			Version:  "v" + version.Version,
			App:      "goradio",
			OS:       runtime.GOOS,
			Arch:     runtime.GOARCH,
			Compiler: runtime.Compiler,
		}

		t, err := template.New("version").Funcs(map[string]any{
			"faint":   lipgloss.NewStyle().Faint(true).Render,
			"bold":    lipgloss.NewStyle().Bold(true).Render,
			"magenta": lipgloss.NewStyle().Foreground(lipgloss.Color("14")).Render,
		}).Parse(`{{ magenta "▇▇▇" }} {{ magenta .App }} 

  {{ faint "Version" }}  {{ bold .Version }}
  {{ faint "Platform" }} {{ bold .OS }}/{{ bold .Arch }}
  {{ faint "Compiler" }} {{ bold .Compiler }}
`)
		handleErr(err)
		handleErr(t.Execute(cmd.OutOrStdout(), versionInfo))
	},
}
