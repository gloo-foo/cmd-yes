package command

import (
	"context"

	gloo "github.com/gloo-foo/framework"
)

// defaultText is the operand GNU yes repeats when none is given.
const defaultText YesText = "y"

// Yes returns a Source that repeatedly emits a line until the downstream
// consumer stops reading (the SIGPIPE analogue), mirroring GNU yes.
//
// Flags:
//   - YesText(s): the operand to repeat (default "y")
//   - YesCount(n): stop after n lines; 0 (the default) repeats forever
func Yes(opts ...any) gloo.Source[[]byte] {
	f := gloo.NewParameters[gloo.File, flags](opts...).Flags
	if f.text == "" {
		f.text = defaultText
	}
	return yesSource{line: []byte(f.text), count: f.count}
}

// yesSource is an immutable Source: its fields are fixed at construction, so a
// value is safe to copy and reuse across pipelines.
type yesSource struct {
	line  []byte
	count YesCount
}

// Stream emits copies of the configured line. Each send is a fresh copy so a
// downstream stage may retain or mutate it without aliasing the next emission.
func (s yesSource) Stream(ctx context.Context) gloo.Stream[[]byte] {
	return gloo.Generate(ctx, func(_ context.Context, send func([]byte) bool, _ func(error)) {
		for i := YesCount(0); s.unbounded() || i < s.count; i++ {
			if !send(s.copy()) {
				return
			}
		}
	})
}

// unbounded reports whether the source repeats forever (no count limit).
func (s yesSource) unbounded() bool { return s.count == 0 }

// copy returns a fresh copy of the line so emissions never alias each other.
func (s yesSource) copy() []byte { return append([]byte(nil), s.line...) }
