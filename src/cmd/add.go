package cmd

import (
	"iwan/src/internal/iwanCore"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new URL in config file",
	Run: func(cmd *cobra.Command, args []string) {
		iwanCore.NewConfigurator().AddUrl(args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
