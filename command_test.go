package command_test

import (
	"context"
	"slices"
	"testing"

	gloo "github.com/gloo-foo/framework"

	command "github.com/gloo-foo/cmd-yes"
)

// collect drains a bounded Source into its emitted lines.
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

func TestYes_DefaultOperandIsY(t *testing.T) {
	// yes | head -n 3 — with no operand, the repeated line is "y".
	got := collect(t, command.Yes(command.YesCount(3)))
	if want := []string{"y", "y", "y"}; !slices.Equal(got, want) {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestYes_TextSetsOperand(t *testing.T) {
	// yes ok | head -n 2 — YesText replaces the repeated operand.
	got := collect(t, command.Yes(command.YesText("ok"), command.YesCount(2)))
	if want := []string{"ok", "ok"}; !slices.Equal(got, want) {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestYes_CountBoundsEmission(t *testing.T) {
	// YesCount caps the number of lines emitted.
	got := collect(t, command.Yes(command.YesCount(5)))
	if len(got) != 5 {
		t.Fatalf("got %d lines, want 5", len(got))
	}
}

func TestYes_EmptyTextFallsBackToDefault(t *testing.T) {
	// An explicit empty operand is indistinguishable from no operand, so the
	// default "y" applies — covering the default-text branch.
	got := collect(t, command.Yes(command.YesText(""), command.YesCount(1)))
	if want := []string{"y"}; !slices.Equal(got, want) {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestYes_UnboundedStopsWhenDownstreamStops(t *testing.T) {
	// With no YesCount, yes repeats forever. The consumer reads a few lines then
	// Discards, which must tear the producer down (the SIGPIPE analogue) — proving
	// the unbounded path and the send-returns-false return branch.
	src := command.Yes(command.YesText("forever"))
	stream := src.Stream(context.Background())
	got := make([]string, 0, 4)
	for item := range stream.Chan() {
		if item.Error != nil {
			t.Fatalf("unexpected error: %v", item.Error)
		}
		got = append(got, string(item.Value))
		if len(got) == 4 {
			stream.Discard()
			break
		}
	}
	if want := []string{"forever", "forever", "forever", "forever"}; !slices.Equal(got, want) {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestYes_EmissionsDoNotAlias(t *testing.T) {
	// Each emitted line must be an independent copy: mutating one must not affect
	// another. This guards the per-send copy.
	ctx := context.Background()
	items, err := gloo.Collect(ctx, command.Yes(command.YesText("ab"), command.YesCount(2)).Stream(ctx))
	if err != nil {
		t.Fatal(err)
	}
	items[0][0] = 'X'
	if string(items[1]) != "ab" {
		t.Fatalf("second emission was aliased: got %q", items[1])
	}
}
