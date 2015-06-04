package robot

import (
	"strings"
	"testing"
)

func TestHelp(t *testing.T) {
	help := Help(&Payload{Text: ""})
	n := strings.Count(help, "\n")
	if n != 2 {
		t.Fatalf("expected %d lines of text", 2)
	}
	t.Logf("%s", help)
}
