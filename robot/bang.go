package robot

type BangBot struct{}

var subcommand map[string]func(*Payload)string
var subdescription map[string]string

func init() { Register("!", &BangBot{}) }

func (b BangBot) Run(p *Payload) string {


	return "fuck you too"
}

func (b BangBot) Description() (string) {
	// return all descriptions
	return "Bang bot!\n\tUsage: !<command>\n\tExpected Response: @user: Pong!"
}

func SubRegister(cmd string, f func(*Payload)string, description string) {
	subcommand[cmd] = f
	subdescription[cmd] = description
}
