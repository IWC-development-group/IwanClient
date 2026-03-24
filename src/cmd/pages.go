package cmd

import (
	"fmt"
	"iwan/src/internal/iwanCore"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

var pagesCmd = &cobra.Command{
	Use:   "pages",
	Short: "Get all the available pages in the selected namespace (Example: iwan pages gl4)",
	Run: func(cmd *cobra.Command, args []string) {
		iwanCore.InitTerminalOutput()

		configurator := iwanCore.NewConfigurator()
		configurator.InitConfig()

		requestedPage := apiPages + os.Args[2]
		response, err := iwanCore.TryAllServers(configurator, requestedPage)
		if err != nil {
			iwanCore.Log("No results")
		}

		renderer, _ := glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
			glamour.WithWordWrap(defaultWidth),
		)

		resPages := "# Namespace: " + response.Namespace + "\n## Available pages:\n"

		for _, value := range response.Pages {
			resPages += "- " + value + "\n"
		}

		result, _ := renderer.Render(resPages)
		fmt.Printf("%s", result)
	},
}

func init() {
	rootCmd.AddCommand(pagesCmd)
}
