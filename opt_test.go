package command_test

import (
	"context"
	"slices"
	"testing"

	gloo "github.com/gloo-foo/framework"

	command "github.com/gloo-foo/cmd-yes"
)

// The yes flag types configure by value and are adapted to the framework's
// Switch seam on the way into NewParameters. Non-flag arguments must pass
// through that adapter untouched so the framework classifies them exactly as
// it always did.

func TestYes_NonFlagArgumentPassesThrough(t *testing.T) {
	// yes ignores positionals: a plain string is classified by the framework as
	// a File positional and never affects the repeated operand — covering the
	// adapter's pass-through branch.
	ctx := context.Background()
	items, err := gloo.Collect(ctx, command.Yes("ignored", command.YesCount(2)).Stream(ctx))
	if err != nil {
		t.Fatal(err)
	}
	got := make([]string, len(items))
	for i, item := range items {
		got[i] = string(item)
	}
	if want := []string{"y", "y"}; !slices.Equal(got, want) {
		t.Fatalf("got %q, want %q", got, want)
	}
}
