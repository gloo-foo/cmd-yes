package yes_test

import (
	"github.com/gloo-foo/framework/patterns"

	command "github.com/gloo-foo/cmd-yes"
)

func ExampleYes_text() {
	// yes hello | head -n 2
	patterns.MustRun(command.Yes(command.YesText("hello"), command.YesCount(2)))
	// Output:
	// hello
	// hello
}
