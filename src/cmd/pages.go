package cmd

import (
	"fmt"
	"iwan/src/internal/iwanCore"
	"iwan/src/internal/iwanFormatting"
	"os"

	"github.com/spf13/cobra"
)

var pagesCmd = &cobra.Command{
	Use:   "pages",
	Short: "Get all the available pages in the selected namespace (Example: iwan pages gl4)",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		iwanCore.InitTerminalOutput()

		configurator := iwanCore.NewConfigurator()
		configurator.InitConfig()

		requestedPage := apiPages + args[0]
		response, err := iwanCore.TryAllServers(configurator, requestedPage, true)
		if err != nil {
			fmt.Println("No results")
			os.Exit(1)
		}

		iwanFormatting.Render(response, iwanFormatting.LIST_FORMAT, iwanFormatting.DEFAULT_RENDER, iwanFormatting.RenderParams{
			Status:    false,
			Name:      false,
			Namespace: false,
			Pages:     true,
		}, iwanFormatting.FormatStringParams{
			HeaderMD:  true,
			ContentMD: false,
			FooterMD:  true,
		})
	},
}

func init() {
	rootCmd.AddCommand(pagesCmd)
}
