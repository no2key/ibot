package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/schema"
	"github.com/miekg/ibot/robot"
)

var (
	port       int
	domainToken map[string]string

	token, name string
)

func main() {
	flag.IntVar(&port, "port", 9999, "port to listen on")
	flag.StringVar(&name, "name", "imp", "the name of the bot")
	flag.StringVar(&token, "token", "", "domain tokens as <domain:token>,<domain:token>")
	flag.Parse()

	if token == "" {
		log.Fatal("need token to start")
	}

	tokens := strings.Split(token, ",")
	domainToken = make(map[string]string)
	for _, s := range tokens {
		tok := strings.Split(s, ":")
		if len(tok) != 2 {
			log.Printf("token needs to be domain:token %s", s)
			continue
		}
		domainToken[tok[0]] = tok[1]
	}
	http.HandleFunc("/", HookHandler)
	log.Printf("starting HTTP server on %d", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

func HookHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	d := schema.NewDecoder()
	command := new(robot.OutgoingWebHook)
	err = d.Decode(command, r.PostForm)
	if err != nil {
		log.Println("couldn't parse post request:", err)
	}
	log.Printf("recieved command: %s from \"%s\"\n", command.Text, command.TeamDomain)

	command.Robot = string(command.Text[0])

	rb := robot.Robot(command.Robot)
	if rb == nil {
		jsonResp(w, "no robot that command yet :(")
		return
	}
	if command.Payload.Token != domainToken[command.TeamDomain] {
		log.Printf("token mismatch, got %s, expected %s", command.Payload.Token, domainToken[command.TeamDomain])
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	jsonResp(w, rb.Run(&command.Payload))
}

func jsonResp(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	resp := map[string]string{"text": msg, "username": name, "icon_emoji": ":imp:"}

	r, err := json.Marshal(resp)
	if err != nil {
		log.Println("couldn't marshal hook response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(r)
}
