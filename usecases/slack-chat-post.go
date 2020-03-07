package usecases

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hf-mush/life-notification/infrastructure"
)

// SlackChatWriteBody request body.
type SlackChatWriteBody struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
	Mrkdwn  bool   `json:"mrkdwn"`
}

// SlackChatPost post chat message.
func SlackChatPost(channel string, message string) error {
	requestBody := SlackChatWriteBody{
		Channel: channel,
		Text:    message,
		Mrkdwn:  false,
	}

	reqJSON, err := json.Marshal(&requestBody)
	if err != nil {
		log.Println(fmt.Sprintf("%v", err.Error()))
		return err
	}

	err = infrastructure.SlackChatWrite(reqJSON)
	if err != nil {
		return err
	}
	return nil
}
