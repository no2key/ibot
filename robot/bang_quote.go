package robot

func init() {
	subRegister("quote", "add or list quotes", Quote,
		`Quote
Synopsis: !quote [quote]

Description: List random quote or add the listed quote to the database`)
}

func Quote(p *Payload) string {
	return "no quotes"
}
