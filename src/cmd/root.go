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
	Short: "Brief",
	Long:  "Long Brief",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {

		iwanCore.InitTerminalOutput()

		configurator := iwanCore.NewConfigurator()
		configurator.InitConfig()

		requestedPage := os.Args[1]

		response, err := iwanCore.TryAllServers(configurator, requestedPage)
		if err != nil {
			iwanCore.Log("No results")
		}

		renderer, _ := glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
			glamour.WithWordWrap(120),
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
