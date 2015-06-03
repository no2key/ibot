package robot

import "strings"

type BangBot struct{}

var subcommand map[string]func(*Payload) string
var subdescription map[string]string
var subshort map[string]string

func init() { Register("!", &BangBot{}) }

func (b BangBot) Run(p *Payload) string {
	// everything up to the first space is the command
	n := strings.Index(p.Text, " ")
	cmd, ok := subcommand[p.Text[:n]]
	if ok {
		p.Text = p.Text[n+1:]
		s := cmd(p)
		return s
	}
	// Check subcommands and run the bot
	return "no command found for " + p.Text
}

func (b BangBot) Description() string {
	// return all descriptions
	return "Bang bot!\n\tUsage: !<command>\n\tExpected Response: @user: Pong!"
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
