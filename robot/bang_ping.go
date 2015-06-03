package robot

func init() {
	SubRegister("ping", "ping your self", Ping,
		`Ping
	Synopsis: !ping

	Description: Ping your self, the expected response will
	be: @user: Pong!`)
}

func Ping(p *Payload) string {
	return "@" + p.UserID + ": pong!"
}
