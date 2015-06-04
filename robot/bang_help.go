package robot

import (
	"bytes"
	"fmt"
	"sort"
)

func init() {
	SubRegister("help", "get help", Help,
		`Help
	Synopsis: !help [command]

	Description: Get a listing of all commands, or specific help on
	the command listed.`)
}

// Sort sorts a map and returns the keys in sorted order.
func Sort(m map[string]string) []string {
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func Help(p *Payload) string {
	if p.Text == "" {
		b := &bytes.Buffer{}
		w := tabWriter(b)
		fmt.Fprintf(w, "COMMAND\tHELP\n")

		keys := Sort(subshort)

		for _, cmd := range keys {
			fmt.Fprintf(w, cmd+"\t"+subshort[cmd]+"\n")
		}
		w.Flush()
		return b.String()
	}
	if v, ok := subdescription[p.Text]; ok {
		return v
	}
	return "no help available for " + p.Text
}
