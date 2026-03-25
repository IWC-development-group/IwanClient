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

		iwanFormatting.Render([]iwanCore.IwanResponse{}, CONFIG_FORMAT, iwanFormatting.DEFAULT_RENDER, iwanFormatting.RenderParams{}, iwanFormatting.FormatStringParams{
			HeaderMD:  true,
			ContentMD: true,
			FooterMD:  true,
		})
	},
}

func CONFIG_FORMAT(content []iwanCore.IwanResponse, params iwanFormatting.RenderParams) []iwanFormatting.FormatString {
	configurator := iwanCore.NewConfigurator()
	configurator.InitConfig()

	header := "# Current URLS:"
	resUrls := ""

	for _, value := range configurator.URLS {
		resUrls += "- " + value + "\n"
	}

	return []iwanFormatting.FormatString{{header, resUrls, ""}}
}

func init() {
	rootCmd.AddCommand(configCmd)
}
