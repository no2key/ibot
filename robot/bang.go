package robot

import (
	"log"
	"strings"
)

type BangBot struct{}

var subcommand map[string]func(*Payload) string
var subdescription map[string]string
var subshort map[string]string

func init() { Register("!", &BangBot{}) }

func (b BangBot) Run(p *Payload) string {
	n := strings.Index(p.Text, " ")
	if n == -1 {
		log.Printf("no command found for: %s", p.Text)
		return "no command found for " + p.Text
	}
	log.Printf("running: %s", p.Text)
	cmd, ok := subcommand[p.Text[:n]]
	if ok {
		p.Text = p.Text[n+1:]
		s := cmd(p)
		return s
	}
	return "no command found for " + p.Text
}

func SubRegister(cmd, short string, f func(*Payload) string, description string) {
	if subcommand == nil {
		subcommand = make(map[string]func(*Payload) string)
		subdescription = make(map[string]string)
		subshort = make(map[string]string)
	}

	subcommand[cmd] = f
	subdescription[cmd] = description
	subshort[cmd] = description
}
