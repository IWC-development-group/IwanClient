package iwanCore

import "fmt"

func Log(content string) {
	if Debug {
		fmt.Println(content)
	}
}
