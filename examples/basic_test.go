package yes_test

import (
	command "github.com/gloo-foo/cmd-yes"
	"github.com/gloo-foo/framework/patterns"
)

func ExampleYes_basic() {
	// yes | head -n 3
	patterns.MustRun(command.Yes(command.YesCount(3)))
	// Output:
	// y
	// y
	// y
}
