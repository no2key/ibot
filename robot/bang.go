package robot


// TODO: !help !commands
type BangBot struct{}

var subcommand map[string]func(*Payload) string
var subdescription map[string]string
var subshort map[string]string

func init() { Register("!", &BangBot{}) }

func (b BangBot) Run(p *Payload) string {
	// Check subcommands and run the bot

	return "fuck you too"
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
