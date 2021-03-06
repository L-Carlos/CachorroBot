package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const (
	telegramAPIBaseURL = "https://api.telegram.org/bot"
	sendMessage        = "/sendMessage"
	tokenEnv           = "TELEGRAM_BOT_TOKEN"
)

var telegramAPI string = telegramAPIBaseURL + os.Getenv(tokenEnv) + sendMessage
var port string = os.Getenv("PORT")

func main() {
	http.HandleFunc("/", handleTelegramWebHook)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleTelegramWebHook(w http.ResponseWriter, r *http.Request) {

	update, err := decodeUpdate(r)

	if err != nil {
		log.Printf("error parsing update, %s\n", err.Error())
		return
	}

	dog, err := getDog()
	if err != nil {
		log.Printf("got error when calling dog API %s\n", err.Error())
		return
	}

	telegramResponseBody, err := sendDogToChat(update.Message.Chat.ID, dog)
	if err != nil {
		log.Printf("got error %s from telegram, response body is %s\n", err.Error(), telegramResponseBody)
	} else {
		log.Printf("Dog %s successfully distributed to chat id %d\n", dog, update.Message.Chat.ID)
	}
}

func sendDogToChat(chatID int, dog string) (string, error) {

	log.Printf("Sending %s to chat_id: %d\n", dog, chatID)
	resp, err := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id": {strconv.Itoa(chatID)},
			"text":    {dog},
		})

	if err != nil {
		log.Printf("error when posting text to the chat: %s\n", err.Error())
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error in parsing telegram answer %s\n", err.Error())
		return "", err
	}
	bodyString := string(bodyBytes)
	log.Printf("Body of Telegram Response: %s\n", bodyString)

	return bodyString, nil
}
