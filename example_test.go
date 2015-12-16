package uiutil_test

import (
	"fmt"

	"github.com/gosuri/uiutil"
)

func ExamplePrompter_PromptHiddenString() {
	var pass string
	uiutil.PromptHiddenString(&pass, "Password: ")
	fmt.Println("Password is", pass)
}
