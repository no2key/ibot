package robot

import (
	"io"
	"text/tabwriter"
)

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

type OutgoingWebHook struct {
	Payload
	TriggerWord string `schema:"trigger_word"`
}

func tabWriter(w io.Writer) *tabwriter.Writer {
	w1 := new(tabwriter.Writer)

	// Format in tab-separated columns with a tab stop of 8.
	w1.Init(w, 0, 8, 1, '\t', 0)
	return w1
}

type Roboter interface {
	Run(p *Payload) string
}

var robots = make(map[string]Roboter)

func Register(command string, r Roboter) {
	robots[command] = r
}

func Robot(command string) Roboter {
	if r, ok := robots[command]; ok {
		return r
	}
	return nil
}
