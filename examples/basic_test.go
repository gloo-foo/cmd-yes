package yes_test

import (
	"github.com/gloo-foo/framework/patterns"

	command "github.com/gloo-foo/cmd-yes"
)

func ExampleYes_basic() {
	// yes | head -n 3
	patterns.MustRun(command.Yes(command.YesCount(3)))
	// Output:
	// y
	// y
	// y
}
