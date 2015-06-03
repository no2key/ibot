package robot

type BangBot struct{}

func init() { Register("!", &BangBot{}) }

func (b BangBot) Run(p *Payload) string {
	return "fuck you too"
}

func (b BangBot) Description() (description string) {
	return "Bang bot!\n\tUsage: !<command>\n\tExpected Response: @user: Pong!"
}
