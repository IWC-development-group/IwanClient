package main

import (
	"fmt"
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
	ip, err := GetWorkingServer(configurator)
	if err != nil {
		panic("Can't found opened server")
	}
	fmt.Println("Founded working ip: " + ip)

	/*requestedPage := os.Args[1]
	response, err := http.Get("http://localhost:8080?name=" + requestedPage)
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
	fmt.Printf("%s (%s)\n---\n%s\n", iwanResponse.Name, iwanResponse.Namespace, result)*/
}
