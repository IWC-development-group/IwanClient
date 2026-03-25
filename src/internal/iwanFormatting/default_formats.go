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

func header_format(content iwanCore.IwanResponse, params RenderParams) string {
	result := ""
	if params.Name && params.Namespace {
		result += content.Name + "(" + content.Namespace + ")"
	} else if params.Name {
		result += content.Name
	}

	if params.Status {
		result += " " + content.Status + "\n"
	}
	return result
}

func DEFAULT_FORMAT(content iwanCore.IwanResponse, params RenderParams) string {
	result := header_format(content, params)
	result += "\n---\n" + content.Content
	return result
}

func LIST_FORMAT(content iwanCore.IwanResponse, params RenderParams) string {
	result := header_format(content, params)

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

	return result
}
