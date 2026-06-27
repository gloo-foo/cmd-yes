package alias_test

import (
	"context"
	"slices"
	"testing"

	yes "github.com/gloo-foo/cmd-yes/alias"
	gloo "github.com/gloo-foo/framework"
)

// The alias package re-exports the constructor and flag types under unprefixed
// names. A mis-wired re-export (say, Count bound to Text, or Yes bound to the
// wrong function) compiles cleanly, so only behavior can prove the wiring. Each
// test exercises one re-export and asserts the yes output it must produce.

// collect runs a bounded Source and returns its emitted lines as strings. Tests
// always supply a Count so the otherwise-infinite stream terminates.
func collect(t *testing.T, src gloo.Source[[]byte]) []string {
	t.Helper()
	ctx := context.Background()
	items, err := gloo.Collect(ctx, src.Stream(ctx))
	if err != nil {
		t.Fatal(err)
	}
	out := make([]string, len(items))
	for i, item := range items {
		out[i] = string(item)
	}
	return out
}

func assertLines(t *testing.T, got, want []string) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestAlias_CountBoundsDefaultLine(t *testing.T) {
	// yes | head -n 3 — the default operand is "y", repeated Count times.
	got := collect(t, yes.Yes(yes.Count(3)))
	assertLines(t, got, []string{"y", "y", "y"})
}

func TestAlias_TextSetsTheOperand(t *testing.T) {
	// yes hello | head -n 2 — Text replaces the repeated operand.
	got := collect(t, yes.Yes(yes.Text("hello"), yes.Count(2)))
	assertLines(t, got, []string{"hello", "hello"})
}
