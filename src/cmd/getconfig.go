package cmd

import (
	"iwan/src/internal/iwanCore"
	"iwan/src/internal/iwanFormatting"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Returns current URLS for check in configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		iwanCore.InitTerminalOutput()

		iwanFormatting.Render(iwanCore.IwanResponse{}, CONFIG_FORMAT, iwanFormatting.DEFAULT_MD_RENDER, iwanFormatting.RenderParams{
			Status:     true,
			Name:       true,
			Namespace:  true,
			Namespaces: true,
		})
	},
}

func CONFIG_FORMAT(content iwanCore.IwanResponse, params iwanFormatting.RenderParams) string {
	configurator := iwanCore.NewConfigurator()
	configurator.InitConfig()

	resUrls := "# Current URLS:\n"

	for _, value := range configurator.URLS {
		resUrls += "- " + value + "\n"
	}

	return resUrls
}

func init() {
	rootCmd.AddCommand(configCmd)
}
