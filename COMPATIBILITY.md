# Yes Command Compatibility

## Summary

Compatible with GNU `yes` for the default and single-operand forms (verified byte-for-byte in the `yupsh/bin-yes` Docker integration harness against Debian coreutils, with both streams capped via `head` since `yes` is infinite). yup-yes adds a non-GNU `-n`/`--count` limit and joins multiple operands with spaces; an explicit empty operand collapses to the default.

## Key Behaviors

```bash
# Default operand: repeat "y" forever (capped here for display)
$ yes | head -n 3
y
y
y

# Single STRING operand: repeat it verbatim
$ yes hello | head -n 3
hello
hello
hello

# -n / --count COUNT: stop after COUNT lines (yup-yes extension; GNU has no such flag)
$ yup-yes -n 3
y
y
y

# Multiple operands: yup-yes joins them with spaces (GNU uses only the first)
$ yup-yes -n 2 one two three
one two three
one two three
```

## Intentional Divergences

- `-n COUNT` / `--count COUNT` is a yup-yes-specific extension with no GNU equivalent: it stops the otherwise-infinite stream after `COUNT` lines. A count of `0` (the default) means repeat forever. GNU `yes` is always infinite and must be bounded externally (e.g. piped to `head`).
- Multiple operands are joined with a single space and the joined string is repeated. GNU `yes one two three` repeats only the first operand (`one`).
- An explicit empty operand collapses to the default `y`: `yup-yes ""` repeats `y`, whereas GNU `yes ""` repeats a blank line. cmd-yes treats an empty `YesText` as "unset" and substitutes the default.
- Apart from the above, the default (`yes`) and single-operand (`yes STRING`) forms are byte-identical to GNU `yes`.
