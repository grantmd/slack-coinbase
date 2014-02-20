package main

// Slack outgoing and incoming webhooks are handled here. Requests come in and
// are examined to see if we need to respond.
//
// Create an outgoing webhook in your Slack here:
// https://my.slack.com/services/new/outgoing-webhook

import (
	"encoding/json"
	"fmt"
	"github.com/grantmd/go-coinbase"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type WebhookResponse struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		incomingText := r.PostFormValue("text")
		if incomingText != "" && r.PostFormValue("user_id") != "" {
			log.Printf("Handling incoming request: %s", incomingText)

			if strings.HasPrefix(incomingText, botUsername) {
				c := &coinbase.Client{}

				rate, err := c.PricesSpotRate()
				if err != nil {
					log.Fatal(err)
				}

				var response WebhookResponse
				response.Username = botUsername
				response.Text = fmt.Sprintf("Price: %f", rate.Amount)
				log.Printf("Sending response: %s", response.Text)

				b, err := json.Marshal(response)
				if err != nil {
					log.Fatal(err)
				}

				w.Write(b)
			}
		}

	})
}

func StartServer(port int) {
	log.Printf("Starting HTTP server on %d", port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
