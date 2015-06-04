package robot

import (
	"strings"
	"testing"
)

func TestHelp(t *testing.T) {
	help := Help(&Payload{Text: ""})
	n := strings.Count(help, "\n")
	if n != len(subcommand) {
		t.Fatalf("expected %d lines of text", len(subcommand))
	}
	t.Logf("%s", help)
}
