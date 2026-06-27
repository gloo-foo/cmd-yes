package command

// YesText sets the string to repeat. Default "y" (GNU yes with no operand).
type YesText string

// YesCount limits the number of repetitions. Zero means repeat forever until
// the downstream consumer stops reading. It is an extension beyond POSIX yes,
// which always repeats indefinitely.
type YesCount int

// Configure records the operand to repeat.
func (t YesText) Configure(flags *flags) { flags.text = t }

// Configure records the repetition limit.
func (c YesCount) Configure(flags *flags) { flags.count = c }

// flags aggregates the parsed yes options.
type flags struct {
	text  YesText
	count YesCount
}
