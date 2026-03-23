package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/charmbracelet/glamour"
)

type IwanResponse struct {
	Status    string `json: "status"`
	Name      string `json: "name"`
	Namespace string `json: "namespace"`
	Content   string `json: "content"`
}

func main() {
	configurator := NewConfigurator()
	configurator.InitConfig()

	//

	initTerminalOutput()

	requestedPage := os.Args[1]
	response, err := http.Get("http://26.70.26.159:8080?name=" + requestedPage)
	if err != nil {
		panic(err.Error())
	}
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}

	var iwanResponse IwanResponse
	unmarshalError := json.Unmarshal(content, &iwanResponse)
	if unmarshalError != nil {
		panic(unmarshalError.Error())
	}

	//XD

	if iwanResponse.Status == "ERR" {
		fmt.Println("Server returned an error: " + iwanResponse.Content)
	}

	renderer, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(80),
	)

	result, _ := renderer.Render(iwanResponse.Content)
	fmt.Printf("%s (%s)\n---\n%s\n", iwanResponse.Name, iwanResponse.Namespace, result)
}
