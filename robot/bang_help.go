package robot

func init() {
	SubRegister("help", "get help", Help,
		`Help
	Synopsis: !help [command]

	Description: Get a listing of all commands, or specific help on
	the command listed.`)
}

func Help(p *Payload) string {
	s := ""
	if p.Text == "" {
		// TODO(miek): sort
		for cmd, help := range subshort {
			s += cmd + " - " + help
		}
		return s
	}
	if v, ok := subdescription[p.Text]; ok {
			return v
	}
	return "no help available for " + p.Text
}
