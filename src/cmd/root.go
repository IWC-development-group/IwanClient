package cmd

import (
	"fmt"
	"iwan/src/internal/iwanCore"
	"iwan/src/internal/iwanFormatting"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "iwan",
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

		response, err := iwanCore.TryAllServers(configurator, requestedPage, false)
		if err != nil {
			fmt.Println("No results")
			os.Exit(1)
		}

		iwanFormatting.Render(response, iwanFormatting.DEFAULT_FORMAT, iwanFormatting.DEFAULT_MD_RENDER, iwanFormatting.RenderParams{
			Status:    true,
			Name:      true,
			Namespace: true,
		})
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
