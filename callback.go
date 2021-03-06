package main

import (
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

func Callback(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Host)
	events, err := bot.ParseRequest(req)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
	}

	for _, event := range events {
		log.Printf("Got event %v", event) //log of event
		log.Println(event.Source.UserID)
		switch event.Type {
		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if err := textHandler(message, event.ReplyToken); err != nil {
					log.Println(err)
				}
			}

		case linebot.EventTypePostback:
			data := event.Postback.Data
			if err := postbackHandler(data); err != nil {
				log.Println(err)
			}
		}
	}
}
