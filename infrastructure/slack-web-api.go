package infrastructure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/shuufujita/life-notification/common"
)

// https://api.slack.com/reference

// SlackChatWriteBody request body.
type SlackChatWriteBody struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
	Mrkdwn  bool   `json:"mrkdwn"`
}

// SlackChatWrite POST request
func SlackChatWrite(message string) error {
	channel := common.GetSlackChannel()
	if channel == "" {
		return fmt.Errorf("Error: %s", "channel is empty")
	}

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

	err = postRequest("chat.postMessage", reqJSON)
	if err != nil {
		log.Println(fmt.Sprintf("error at SlackChatWrite : %v", err.Error()))
		return err
	}
	return nil
}

func postRequest(method string, requestBody []byte) error {
	requestURL := common.GetSlackAPIURL() + "/" + method
	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println(fmt.Sprintf("%v", err.Error()))
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+common.GetSlackAPIToken())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(fmt.Sprintf("%v", err.Error()))
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if resp.StatusCode >= 400 {
		log.Println(fmt.Sprintf("%v %v", resp.StatusCode, string(body)))
		return fmt.Errorf("Error: %s", "not success")
	}
	return nil
}
