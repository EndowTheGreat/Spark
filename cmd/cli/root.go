package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "spark",
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func Setup() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
