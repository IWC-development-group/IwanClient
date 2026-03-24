package cmd

import (
	"fmt"
	"iwan/src/internal/iwanCore"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

var namespacesCmd = &cobra.Command{
	Use:   "namespaces",
	Short: "Brief",
	Long:  "Long Brief",
	Run: func(cmd *cobra.Command, args []string) {
		iwanCore.InitTerminalOutput()

		configurator := iwanCore.NewConfigurator()
		configurator.InitConfig()

		requestedPage := apiNamespaces

		response, err := iwanCore.TryAllServers(configurator, requestedPage)
		if err != nil {
			iwanCore.Log("No results")
		}

		renderer, _ := glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
			glamour.WithWordWrap(120),
		)

		resNamespaces := ""

		for _, value := range response.Namespaces {
			resNamespaces += value + "\n"
		}

		result, _ := renderer.Render(resNamespaces)
		fmt.Printf("%s (%s)\n---\n%s\n", response.Name, response.Namespace, result)
	},
}

func init() {
	rootCmd.AddCommand(namespacesCmd)
}
