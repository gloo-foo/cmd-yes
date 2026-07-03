// Package alias provides unprefixed re-exports for the yes command.
//
//	import yes "github.com/gloo-foo/cmd-yes/alias"
//	yes.Yes(yes.Text("hello"), yes.Count(3))
package alias

import (
	gloo "github.com/gloo-foo/framework"

	command "github.com/gloo-foo/cmd-yes"
)

// Yes re-exports the constructor.
func Yes(opts ...any) gloo.Source[[]byte] { return command.Yes(opts...) }

// Text sets the operand to repeat (default "y").
type Text = command.YesText

// Count limits the number of repetitions (0 = forever).
type Count = command.YesCount
