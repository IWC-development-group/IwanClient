package cmd

import (
	"fmt"
	"iwan/src/internal/iwanCore"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "IwanClient",
	Short: "Search for documentation via namespace/name (Example: iwan gl4/glBindBuffer)",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Welcome to IwanClient! Type --help for usage")
			os.Exit(0)
		}

		iwanCore.InitTerminalOutput()

		configurator := iwanCore.NewConfigurator()
		configurator.InitConfig()

		requestedPage := apiSearch + args[0]

		response, err := iwanCore.TryAllServers(configurator, requestedPage)
		if err != nil {
			iwanCore.Log("No results")
		}

		renderer, _ := glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
			glamour.WithWordWrap(defaultWidth),
		)

		result, _ := renderer.Render(response.Content)
		fmt.Printf("%s (%s)\n---\n%s\n", response.Name, response.Namespace, result)
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&iwanCore.Debug, "debug", "d", false, "Turn on console debugging")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
