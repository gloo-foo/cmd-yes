package yes_test

import (
	"github.com/gloo-foo/framework/patterns"

	command "github.com/gloo-foo/cmd-yes"
)

func ExampleYes_text() {
	// yes hello | head -n 2
	if err := patterns.Run(command.Yes(command.YesText("hello"), command.YesCount(2))); err != nil {
		panic(err)
	}
	// Output:
	// hello
	// hello
}
