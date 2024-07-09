package markdown

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gitlab.com/EndowTheGreat/spark/asset"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

var Output string

func Convert(dir string, extensions parser.Extensions, opts html.RendererOptions) {
	assets := make(map[string]string)
	if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		relPath, err := filepath.Rel(dir, path)
		if err != nil {
			return fmt.Errorf("Failed to determine relative path: %w", err)
		}
		outDir := filepath.Join(Output, filepath.Dir(relPath))
		if err := os.MkdirAll(outDir, 0755); err != nil {
			return fmt.Errorf("Failed to create output directory: %w", err)
		}
		fileName := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
		switch ext := filepath.Ext(path); ext {
		case ".md":
			data, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("Failed to read file: %w", err)
			}
			file := filepath.Join(outDir, fileName+".html")
			if err := os.WriteFile(file, markdown.ToHTML(data, parser.NewWithExtensions(extensions), html.NewRenderer(opts)), 0644); err != nil {
				return fmt.Errorf("Failed to write file: %w", err)
			}
			if css, ok := assets[fileName+".css"]; ok {
				asset.Inject(file, Output, css, "style")
			}
			if js, ok := assets[fileName+".js"]; ok {
				asset.Inject(file, Output, js, "script")
			}
		case ".css", ".js":
			file := filepath.Join(outDir, filepath.Base(path))
			if _, err := asset.Copy(path, file); err != nil {
				return fmt.Errorf("Failed to copy file: %w", err)
			}
			assets[filepath.Base(path)] = file
		default:
			file := filepath.Join(outDir, filepath.Base(path))
			if _, err := asset.Copy(path, file); err != nil {
				return fmt.Errorf("Failed to copy file: %w", err)
			}
		}
		return nil
	}); err != nil {
		log.Fatalf("Error walking the directory: %v", err)
	}
	fmt.Printf("Successfully converted files into the %v directory.\n", Output)
}
