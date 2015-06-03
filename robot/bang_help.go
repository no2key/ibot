package robot

import (
	"bytes"
	"fmt"
)

func init() {
	SubRegister("help", "get help", Help,
		`Help
	Synopsis: !help [command]

	Description: Get a listing of all commands, or specific help on
	the command listed.`)
}

func Help(p *Payload) string {
	if p.Text == "" {
		b := &bytes.Buffer{}
		w := tabWriter(b)
		fmt.Fprintf(w, "COMMAND\tHELP\n")
		for cmd, help := range subshort {
			fmt.Fprintf(w, cmd+"\t"+help+"\n")
		}
		return b.String()
	}
	if v, ok := subdescription[p.Text]; ok {
		return v
	}
	return "no help available for " + p.Text
}
