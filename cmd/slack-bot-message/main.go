package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shuufujita/life-notification/common"
	"github.com/shuufujita/life-notification/usecases"
)

func init() {
	common.LoadDotEnv()
	err := common.CustomLogger()
	if err != nil {
		os.Exit(1)
	}
}

func main() {
	message := "sample message text"
	err := usecases.SlackChatPost(message)
	if err != nil {
		log.Println(fmt.Sprintf("%v: %v", "error", err.Error()))
		os.Exit(1)
		return
	}
	os.Exit(0)
}
