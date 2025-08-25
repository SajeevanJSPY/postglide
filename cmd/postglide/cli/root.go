package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "postglide [cmd]",
	Short: "Postglide is protocol-level router and connection orchestrator for PostgreSQL",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	rootCmd.AddCommand(startCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
