package cmd

import (
	_ "embed"
	"os"

	"github.com/spf13/cobra"
)

//go:embed templates/record.md
var template []byte

var defaultDir = "docs/adr"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "adr",
	Short: "A CLI tool for managing Architecture Decision Records (ADRs).",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(NewAddCmd(template, defaultDir))
	rootCmd.AddCommand(NewInitCmd(template, defaultDir))
}
