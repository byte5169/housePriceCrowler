package utils

import (
	"fmt"
	"net/http"
	"strings"
)

// place your own bot token and chat_ids here
var botToken = "###"
var chatIds = []string{"###", "##"}

// It takes a string as an argument, and sends it to all the chat IDs in the chatIds array
//
// Args:
//   message (string): The message you want to send.
func SendPM(message string) {

	url := fmt.Sprintf("https://api.telegram.org/bot%v/sendMessage", botToken)

	for _, chatId := range chatIds {
		payloadString := fmt.Sprintf("{\"text\":\"%v\",\"parse_mode\":\"HTML\",\"disable_web_page_preview\":false,\"disable_notification\":false,\"reply_to_message_id\":null,\"chat_id\":\"%v\"}", message, chatId)
		payload := strings.NewReader(payloadString)

		req, _ := http.NewRequest("POST", url, payload)

		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")

		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()
	}

}
