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
	p.Text = p.Text[1:]
	cmd := ""

	n := strings.Index(p.Text, " ")
	if n != -1 {
		cmd = p.Text[:n]
		// TODO(miek): can overflow. At least check/test
		p.Text = p.Text[n+1:]
	} else {
		cmd = p.Text
		p.Text = ""
	}

	log.Printf("running: %s with %s", cmd, p.Text)
	c, ok := subcommand[cmd]
	if ok {
		s := c(p)
		return s
	}
	return "no command found for " + cmd
}

// SubRegister registers a subcommand for the bang bot.
func SubRegister(cmd, short string, f func(*Payload) string, description string) {
	if subcommand == nil {
		subcommand = make(map[string]func(*Payload) string)
		subdescription = make(map[string]string)
		subshort = make(map[string]string)
	}

	subcommand[cmd] = f
	subdescription[cmd] = description
	subshort[cmd] = short
}
