package robots

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
)

type SlashCommand struct {
	Payload
	Command string `schema:"command"`
}

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

type Configuration struct {
	Port         int               `json:"port"`
	DomainTokens map[string]string `json:"domain_tokens"`
}

type Robot interface {
	Run(p *Payload) (botString string)
	Description() (description string)
}

var Robots = make(map[string]Robot)
var Config = new(Configuration)
var ConfigDirectory = flag.String("c", ".", "Configuration directory (default .)")

func init() {
	flag.Parse()
	configFile := filepath.Join(*ConfigDirectory, "config.json")
	config, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal("Error opening config: ", err)
	}

	err = json.Unmarshal(config, Config)
	if err != nil {
		log.Fatal("Error parsing config: ", err)
	}

	log.Printf("Found %d domain configurations", len(Config.DomainTokens))
	log.Printf("Port: %d", Config.Port)
}

func RegisterRobot(command string, r Robot) {
	if _, ok := Robots[command]; ok {
		log.Printf("There are two robots mapped to %s!", command)
	} else {
		log.Printf("Registered: %s", command)
		Robots[command] = r
	}
}

func (i *IncomingWebhook) Send() error {
	if _, ok := Config.DomainTokens[i.Domain]; !ok {
		err := errors.New(fmt.Sprintf("Domain: \"%s\" not found in configuration", i.Domain))
		log.Print(err)
		return err
	}

	webhook := url.URL{
		Scheme: "https",
		Host:   i.Domain + ".slack.com",
		Path:   "/services/hooks/incoming-webhook",
	}

	p, err := json.Marshal(i)
	if err != nil {
		return err
	}

	data := url.Values{}
	data.Set("payload", string(p))
	data.Set("token", Config.DomainTokens[i.Domain])

	webhook.RawQuery = data.Encode()
	resp, err := http.PostForm(webhook.String(), data)
	if resp.StatusCode != 200 {
		message := fmt.Sprintf("ERROR: Non-200 Response from Slack Incoming Webhook API: %s", resp.Status)
		log.Println(message)
	}
	return err
}
