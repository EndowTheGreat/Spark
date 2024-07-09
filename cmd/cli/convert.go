package cli

import (
	"os"

	"gitlab.com/EndowTheGreat/spark/markdown"

	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/spf13/cobra"
)

var input string

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().StringVarP(&input, "input", "i", ".", "The directory to read markdown files from.")
	convertCmd.Flags().StringVarP(&markdown.Output, "output", "o", "output", "The directory to write converted HTML files to.")
}

var convertCmd = &cobra.Command{
	Use:     "convert",
	Short:   "Convert markdown files in your input directory to HTML in your output.",
	Aliases: []string{"c"},
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(markdown.Output); os.IsNotExist(err) {
			os.Mkdir(markdown.Output, 0755)
		}
		extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
		opts := html.RendererOptions{Flags: html.CommonFlags | html.HrefTargetBlank}
		markdown.Convert(input, extensions, opts)
	},
}
