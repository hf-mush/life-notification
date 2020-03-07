package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hf-mush/life-notification/common"
	"github.com/hf-mush/life-notification/usecases"
)

func main() {
	common.LoadEnv()

	message := "sample message text"

	err := usecases.SlackChatPost(message)
	if err != nil {
		log.Println(fmt.Sprintf("%v", err.Error()))
		os.Exit(1)
		return
	}
	os.Exit(0)
}
