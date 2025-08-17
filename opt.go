package command

type Count int

type flags struct {
	Count Count
}

func (c Count) Configure(flags *flags) { flags.Count = c }
