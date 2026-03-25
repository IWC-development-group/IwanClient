package cmd

import (
	"iwan/src/internal/iwanCore"
	"iwan/src/internal/iwanFormatting"

	"github.com/spf13/cobra"
)

var namespacesCmd = &cobra.Command{
	Use:   "namespaces",
	Short: "Get all available namespaces on the server",
	Run: func(cmd *cobra.Command, args []string) {
		iwanCore.InitTerminalOutput()

		configurator := iwanCore.NewConfigurator()
		configurator.InitConfig()

		requestedPage := apiNamespaces

		response, err := iwanCore.TryAllServers(configurator, requestedPage)
		if err != nil {
			iwanCore.Log("No results")
		}

		iwanFormatting.Render(response, iwanFormatting.LIST_FORMAT, iwanFormatting.DEFAULT_MD_RENDER, iwanFormatting.RenderParams{
			Status:     false,
			Name:       false,
			Namespace:  false,
			Namespaces: true,
		})
	},
}

func init() {
	rootCmd.AddCommand(namespacesCmd)
}
