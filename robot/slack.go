package robot

import (
	"flag"
	"log"
	"strings"
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

type OutgoingWebHookResponse struct {
	Text      string     `json:"text"`
	Parse     ParseStyle `json:"parse,omitempty"`
	LinkNames bool       `json:"link_names,omitempty"`
	Markdown  bool       `json:"mrkdwn,omitempty"`
}

type ParseStyle string

var (
	ParseStyleFull = ParseStyle("full")
	ParseStyleNone = ParseStyle("none")
)

type IncomingWebhook struct {
	Domain      string       `json:"domain"`
	Channel     string       `json:"channel"`
	Username    string       `json:"username"`
	Text        string       `json:"text"`
	IconEmoji   string       `json:"icon_emoji,omitempty"`
	IconURL     string       `json:"icon_url,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	UnfurlLinks bool         `json:"unfurl_links,omitempty"`
	Parse       ParseStyle   `json:"parse,omitempty"`
	LinkNames   bool         `json:"link_names,omitempty"`
	Markdown    bool         `json:"mrkdwn,omitempty"`
}

type Attachment struct {
	Fallback   string            `json:"fallback"`
	Pretext    string            `json:"pretext,omitempty"`
	Text       string            `json:"text,omitempty"`
	Color      string            `json:"color,omitempty"`
	Fields     []AttachmentField `json:"fields,omitempty"`
	MarkdownIn []MarkdownField   `json:"mrkdown_in,omitempty"`
}

type MarkdownField string

var (
	MarkdownFieldPretext  = MarkdownField("pretext")
	MarkdownFieldText     = MarkdownField("text")
	MarkdownFieldTitle    = MarkdownField("title")
	MarkdownFieldFields   = MarkdownField("fields")
	MarkdownFieldFallback = MarkdownField("fallback")
)

type AttachmentField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short,omitempty"`
}

type Robot interface {
	Run(p *Payload) (botString string)
	Description() (description string)
}

var Robots = make(map[string]Robot)

var (
	port       int
	domainToen map[string]string

	token string
)

func init() {
	flag.IntVar(&port, "port", 9999, "port to listen on")
	flag.StringVar(&token, "token", "", "domain tokens as <domain:token>,<domain:token>")
	flag.Parse()

	tokens := strings.Split(token, ",")
	for _, s := range tokens {
		println(s)
	}
}

func RegisterRobot(command string, r Robot) {
	if _, ok := Robots[command]; ok {
		log.Printf("There are two robots mapped to %s!", command)
	} else {
		log.Printf("Registered: %s", command)
		Robots[command] = r
	}
}
