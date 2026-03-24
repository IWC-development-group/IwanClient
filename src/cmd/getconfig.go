package cmd

import (
	"fmt"
	"iwan/src/internal/iwanCore"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Returns current URLS for check in configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		iwanCore.InitTerminalOutput()

		configurator := iwanCore.NewConfigurator()
		configurator.InitConfig()

		renderer, _ := glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
			glamour.WithWordWrap(defaultWidth),
		)

		resUrls := "# Current URLS:\n"

		for _, value := range configurator.URLS {
			resUrls += "- " + value + "\n"
		}

		result, _ := renderer.Render(resUrls)
		fmt.Printf("%s", result)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
