package robots

type PingBot struct{}

func init() {
	p := &PingBot{}
	RegisterRobot("!", p)
}

func (pb PingBot) Run(p *Payload) string {
	return "fuck you too"
}

func (pb PingBot) Description() (description string) {
	return "Ping bot!\n\tUsage: /ping\n\tExpected Response: @user: Pong!"
}
