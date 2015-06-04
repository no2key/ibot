package robot

import (
	"strings"
	"testing"
)

func TestHelp(t *testing.T) {
	help := Help(&Payload{Text: ""})
	n := strings.Index(help, "\n")
	if n == -1 {
		t.Fatal("help string must contain a newline")
	}
	if help[:n] != "COMMAND		HELP" {
		t.Fatal("COMMAND and HELP titles not found")
	}
}
