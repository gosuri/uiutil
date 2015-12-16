package main

import (
	"fmt"

	"github.com/gosuri/uiutil"
)

func main() {
	var pass string
	uiutil.PromptHiddenString(&pass, "Password: ")
	fmt.Println("Password is", pass)
}
