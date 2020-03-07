package infrastructure

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hf-mush/life-notification/common"
)

// https://api.slack.com/reference

// SlackChatWrite POST request
func SlackChatWrite(requestBody []byte) error {
	err := postRequest("chat.postMessage", requestBody)
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
