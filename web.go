package main

// Slack outgoing and incoming webhooks are handled here. Requests come in and
// are examined to see if we need to respond.
//
// Create an outgoing webhook in your Slack here:
// https://my.slack.com/services/new/outgoing-webhook

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type WebhookResponse struct {
	Username string `json:"username"`
	Text     string `json:"text"`
	Channel  string `json:"channel"`
}

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		incomingText := r.PostFormValue("text")
		log.Printf("Handling incoming request: %s", incomingText)

	})
}

func StartServer(port int) {
	log.Printf("Starting HTTP server on %d", port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
