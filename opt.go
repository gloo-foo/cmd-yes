package command

// YesText sets the string to repeat. Default "y" (GNU yes with no operand).
type YesText string

// YesCount limits the number of repetitions. Zero means repeat forever until
// the downstream consumer stops reading. It is an extension beyond POSIX yes,
// which always repeats indefinitely.
type YesCount int

// configure records the operand to repeat.
func (t YesText) configure(f flags) flags {
	f.text = t
	return f
}

// configure records the repetition limit.
func (c YesCount) configure(f flags) flags {
	f.count = c
	return f
}

// flags aggregates the parsed yes options.
type flags struct {
	text  YesText
	count YesCount
}

// valueFlag is a yes option that configures the flags immutably — value in,
// value out — so the package's own flag types never mutate through a pointer.
type valueFlag interface {
	configure(flags) flags
}

// option adapts a value-semantic flag to the framework's Switch seam. The
// pointer in Configure is the framework's generic contract (Switch[T] requires
// Configure(*T)); it never escapes this adapter.
type option[T any] func(T) T

// Configure implements gloo.Switch[T] by applying the value-semantic mutation
// through the framework's pointer seam.
func (o option[T]) Configure(t *T) { *t = o(*t) }

// options adapts yes's value-semantic flags to the framework Switch contract,
// passing every other argument through untouched.
func options(args []any) []any {
	adapted := make([]any, len(args))
	for i, arg := range args {
		adapted[i] = adapt(arg)
	}
	return adapted
}

// adapt wraps a single value-semantic flag as a framework Switch.
func adapt(arg any) any {
	if f, ok := arg.(valueFlag); ok {
		return option[flags](f.configure)
	}
	return arg
}
