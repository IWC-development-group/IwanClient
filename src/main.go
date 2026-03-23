package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/glamour"
	//"github.com/spf13/cobra"
)

type IwanResponse struct {
	Status    string `json: "status"`
	Name      string `json: "name"`
	Namespace string `json: "namespace"`
	Content   string `json: "content"`
}

func main() {
	initTerminalOutput()

	configurator := NewConfigurator()
	configurator.InitConfig()

	requestedPage := os.Args[1]

	response, err := TryAllServers(configurator, requestedPage)
	if err != nil {
		panic("No results")
	}

	renderer, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(120),
	)

	result, _ := renderer.Render(response.Content)
	fmt.Printf("%s (%s)\n---\n%s\n", response.Name, response.Namespace, result)
}
