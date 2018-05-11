package ciscosparkbot

import (
	"fmt"
	"log"

	"github.com/parnurzeal/gorequest"
)

type Bot struct {
	Name        string
	Accesstoken string
}

type Bots struct {
	Bot []Bot
}

type Message struct {
	RoomID string `json:"roomId"`
	Text   string `json:"text"`
}

var hostname string

func (b *Bot) Sendmessage(roomid string, msg string) {
	url := "https://api.ciscospark.com/v1/messages"
	token := "Bearer " + b.Accesstoken
	res, _, _ := gorequest.New().Post(url).SendStruct(Message{RoomID: roomid, Text: msg}).Set("Content-Type", "application/json").Set("Authorization", token).End()

	log.Println(res)

}

func (b *Bot) Getmessage(msgid string) string {
	//https://developer.webex.com/endpoint-messages-messageId-get.html
	//https://api.ciscospark.com/v1/messages/msgid
	//https://api.ciscospark.com/v1/messages/Y2lzY29zcGFyazovL3VzL01FU1NBR0UvYzk2YjI3ZDAtNTQyYi0xMWU4LTgyMDItY2Q4OGE4ZDNhYjQ3
	url := "https://api.ciscospark.com/v1/messages/" + msgid
	token := "Bearer " + b.Accesstoken
	fmt.Println(url, token)
	_, body, _ := gorequest.New().Get(url).Set("Content-Type", "application/json").Set("Authorization", token).End()
	fmt.Println(body)
	return body

}

func (b *Bot) Registermessagewebhook() {
	//https://developer.webex.com/endpoint-webhooks-post.html
	url := "https://api.ciscospark.com/v1/webhooks"
	token := "Bearer " + b.Accesstoken
	_, body, _ := gorequest.New().Post(url).Set("Content-Type", "application/json").Set("Authorization", token).
		Send(`{"name":"jirabot", "targetUrl":"http://478cbd16.ngrok.io/jirawebhook","resource":"messages","event":"created"}`).End()
	log.Println(body)
}

func Newbot(accesstoken string) *Bot {
	b := Bot{Accesstoken: accesstoken}
	return &b
}
