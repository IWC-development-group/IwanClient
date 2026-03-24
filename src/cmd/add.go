package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Brief",
	Long:  "Long Brief",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Тут типа поиск будет")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
