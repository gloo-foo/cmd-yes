// Package alias provides unprefixed re-exports for the yes command.
//
//	import yes "github.com/gloo-foo/cmd-yes/alias"
//	yes.Yes(yes.Text("hello"), yes.Count(3))
package alias

import command "github.com/gloo-foo/cmd-yes"

// Yes re-exports the constructor.
var Yes = command.Yes

// Text sets the operand to repeat (default "y").
type Text = command.YesText

// Count limits the number of repetitions (0 = forever).
type Count = command.YesCount
