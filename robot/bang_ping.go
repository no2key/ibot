package robot

func init() {
	SubRegister("ping", "ping a user", Ping,
		`Ping bot!
	Usage: !ping
	Expected Response: @user: Pong!`)
}

func Ping(p *Payload) string {
	return "@user: pong!"
}
