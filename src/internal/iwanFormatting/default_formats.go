package iwanFormatting

import (
	"fmt"
	"iwan/src/internal/iwanCore"
	"strings"

	"github.com/charmbracelet/glamour"
)

//Render

func DEFAULT_RENDER(contents []FormatString, params FormatStringParams) {
	renderer, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(defaultWidth),
	)

	for _, content := range contents {
		resultHeader, _ := renderer.Render(content.Header)
		resultContent, _ := renderer.Render(content.Content)
		resultFooter, _ := renderer.Render(content.Footer)

		if !params.HeaderMD {
			resultHeader = content.Header
		}
		if !params.ContentMD {
			resultContent = content.Content
		}
		if !params.FooterMD {
			resultFooter = content.Footer
		}

		fmt.Printf("%s\n%s\n%s", resultHeader, resultContent, resultFooter)
	}
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
			result += " " + content.Status
		}
	}
	return result
}

func DEFAULT_FORMAT(contents []iwanCore.IwanResponse, params RenderParams) []FormatString {
	header := header_format(contents, params)
	result := ""
	footer := ""

	for _, content := range contents {
		result += "\n---\n" + content.Content + "\n"
		footer = "## Retrieved from: " + content.Address
	}

	return []FormatString{{header, result, footer}}
}

func LIST_FORMAT(contents []iwanCore.IwanResponse, params RenderParams) []FormatString {
	outList := []FormatString{}

	var columnCount = 1

	for _, content := range contents {
		header := ""
		result := ""
		footer := ""

		if params.Namespaces {
			var sb strings.Builder
			header = "# Available namespaces:"

			calc_column(&columnCount, len(content.Namespaces))

			for i, value := range content.Namespaces {
				item := fmt.Sprintf("%-25s", value)
				sb.WriteString(item)

				if (i+1)%columnCount == 0 && i != len(content.Namespaces)-1 {
					sb.WriteString("\n")
				}
				//result += "- " + value + "\n"
			}
			result += sb.String()
		}

		if params.Pages {
			var sb strings.Builder
			header = "# Namespace: " + content.Namespace + "\n## Available pages:"

			calc_column(&columnCount, len(content.Pages))

			for i, value := range content.Pages {
				item := fmt.Sprintf("%-25s", value)
				sb.WriteString(item)

				if (i+1)%columnCount == 0 && i != len(content.Pages)-1 {
					sb.WriteString("\n")
				}
			}
			result += sb.String()
		}
		footer = "## Retrieved from: " + content.Address
		outList = append(outList, FormatString{header, result, footer})
	}
	//fmt.Println(len(outList))
	return outList
}

func calc_column(col *int, count int) {
	//fmt.Println("count: " + strconv.Itoa(count))
	if count <= 30 {
		*col = 1
	} else if count >= 30 && count <= 100 {
		*col = 2
	} else {
		*col = 3
	}
	//fmt.Println("size: " + strconv.Itoa(*col))
}
