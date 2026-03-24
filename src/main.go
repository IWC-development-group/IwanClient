package main

import "iwan/src/cmd"

//"github.com/charmbracelet/glamour"
//"github.com/spf13/cobra"

func main() {
	cmd.Execute()

	/*
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
		fmt.Printf("%s (%s)\n---\n%s\n", response.Name, response.Namespace, result)*/
}
