package robots

type BangBot struct{}

func init() {
	b := &BangBot{}
	RegisterRobot("!", b)
}

func (b BangBot) Run(p *Payload) string {
	return "fuck you too"
}

func (b BangBot) Description() (description string) {
	return "Bang bot!\n\tUsage: !<command>\n\tExpected Response: @user: Pong!"
}
