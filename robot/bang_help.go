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
		s = "COMMAND\t\tHELP\n"
		for cmd, help := range subshort {
			s += cmd + "\t\t" + help + "\n"
		}
		return s
	}
	if v, ok := subdescription[p.Text]; ok {
			return v
	}
	return "no help available for " + p.Text
}
