package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Update is a Telegram object received by the handler
// when a user interacts with the bot
type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

// Message is a Telegram object in Update
type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

// Chat is a Telegram object in Message to identify the chat
type Chat struct {
	ID int `json:"id"`
}

func decodeUpdate(r *http.Request) (*Update, error) {
	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("Could not decode update: %s\n", err.Error())
		return nil, err
	}
	return &update, nil
}
