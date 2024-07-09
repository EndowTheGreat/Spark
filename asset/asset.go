package asset

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Copy(src, dst string) (string, error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return "", fmt.Errorf("could not open source file: %w", err)
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dst)
	if err != nil {
		return "", fmt.Errorf("could not create destination file: %w", err)
	}
	defer dstFile.Close()
	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return "", fmt.Errorf("could not copy file content: %w", err)
	}
	return dst, nil
}

func Inject(output string, outDir string, assetFile, tag string) {
	var asset string
	file, err := os.OpenFile(output, os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("Could not open output file: %v", err)
	}
	defer file.Close()
	document, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatalf("Could not parse HTML document: %v", err)
	}
	switch tag {
	case "style":
		asset = fmt.Sprintf("<link rel=\"stylesheet\" href=\"%s\"/>", strings.TrimPrefix(assetFile, outDir))
		document.Find("head").AppendHtml(asset)
	case "script":
		content, err := os.ReadFile(assetFile)
		if err != nil {
			log.Fatalf("Could not read asset file: %v", err)
		}
		asset = fmt.Sprintf("<script>\n%s\n</script>", content)
		document.Find("body").AppendHtml(asset)
	}
	content, err := document.Html()
	if err != nil {
		log.Fatalf("Could not get HTML from document: %v", err)
	}
	if err := os.WriteFile(output, []byte(content), 0644); err != nil {
		log.Fatalf("Could not save modified HTML document: %v", err)
	}
}
