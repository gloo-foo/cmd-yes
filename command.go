package command

import (
	"context"
	"io"

	yup "github.com/gloo-foo/framework"
)

type command yup.Inputs[string, flags]

func Yes(parameters ...any) yup.Command {
	return command(yup.Initialize[string, flags](parameters...))
}

func (p command) Executor() yup.CommandExecutor {
	return func(ctx context.Context, stdin io.Reader, stdout, stderr io.Writer) error {
		// Determine the string to output
		output := "y"
		if len(p.Positional) > 0 {
			// Join all positional arguments with spaces
			output = p.Positional[0]
			for i := 1; i < len(p.Positional); i++ {
				output += " " + p.Positional[i]
			}
		}
		output += "\n"

		// Determine repetition count
		count := int(p.Flags.Count)
		infinite := count <= 0

		// Output loop
		for i := 0; infinite || i < count; i++ {
			// Check for context cancellation
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}

			// Write output
			_, err := io.WriteString(stdout, output)
			if err != nil {
				return err
			}
		}

		return nil
	}
}
