package iwanFormatting

import (
	"fmt"
	"iwan/src/internal/iwanCore"

	"github.com/charmbracelet/glamour"
)

//Render

func DEFAULT_MD_RENDER(content string) {
	renderer, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(defaultWidth),
	)

	result, _ := renderer.Render(content)
	fmt.Printf("%s", result)
}

//Format

func header_format(contents []iwanCore.IwanResponse, params RenderParams) string {
	result := ""

	for _, content := range contents {
		if params.Name && params.Namespace {
			result += content.Name + "(" + content.Namespace + ")"
		} else if params.Name {
			result += content.Name
		}

		if params.Status {
			result += " " + content.Status + "\n"
		}
	}
	return result
}

func DEFAULT_FORMAT(contents []iwanCore.IwanResponse, params RenderParams) string {
	result := header_format(contents, params)
	for _, content := range contents {
		result += "\n---\n" + content.Content + "\n"
		result += "## Answer from: " + content.Address + "\n"
	}
	return result
}

func LIST_FORMAT(contents []iwanCore.IwanResponse, params RenderParams) string {
	result := header_format(contents, params)

	for _, content := range contents {
		if params.Namespaces {
			result += "# Available namespaces:\n"
			for _, value := range content.Namespaces {
				result += "- " + value + "\n"
			}
		}

		if params.Pages {
			result += "# Namespace: " + content.Namespace + "\n## Available pages:\n"
			for _, value := range content.Pages {
				result += "- " + value + "\n"
			}
		}
		result += "## Answer from: " + content.Address + "\n"
	}

	return result
}
