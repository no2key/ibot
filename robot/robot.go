package robot

import (
	"io"
	"text/tabwriter"
)

// Payload is the payload Slack sends a bot.
type Payload struct {
	Token       string  `schema:"token"`
	TeamID      string  `schema:"team_id"`
	TeamDomain  string  `schema:"team_domain,omitempty"`
	ChannelID   string  `schema:"channel_id"`
	ChannelName string  `schema:"channel_name"`
	Timestamp   float64 `schema:"timestamp,omitempty"`
	UserID      string  `schema:"user_id"`
	UserName    string  `schema:"user_name"`
	Text        string  `schema:"text,omitempty"`
	TriggerWord string  `schema:"trigger_word,omitempty"`
	Service_ID  string  `schema:"service_id,omitempty"`
	Robot       string
}

// OutgoingWebHook is Slack's Outgoing Webhook.
type OutgoingWebHook struct {
	Payload
	TriggerWord string `schema:"trigger_word"`
}

func tabWriter(w io.Writer) *tabwriter.Writer {
	w1 := new(tabwriter.Writer)
	// Format in tab-separated columns with a tab stop of 8.
	w1.Init(w, 0, 8, 2, '\t', 0)
	return w1
}

// Roboter is an interface that a Bot must implement.
type Roboter interface {
	Run(p *Payload) string
}

var robots = make(map[string]Roboter)

// Register registers a bot.
func Register(command string, r Roboter) {
	robots[command] = r
}

// Robot returns a registered bot or nil.
func Robot(command string) Roboter {
	if r, ok := robots[command]; ok {
		return r
	}
	return nil
}
