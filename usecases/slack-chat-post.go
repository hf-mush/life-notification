package usecases

import "github.com/hf-mush/life-notification/infrastructure"

// SlackChatPost post chat message.
func SlackChatPost(message string) error {
	// TODO: message custmize / validation
	err := infrastructure.SlackChatWrite(message)
	if err != nil {
		return err
	}
	return nil
}
