package handler

import (
	"log"

	keys "github.com/RyuseiNomi/shinkansen_bot/keys"
	text "github.com/RyuseiNomi/shinkansen_bot/text"
)

func Tweet() {
	log.Print("aaa")
	text := text.GetTweetText()
	api := keys.GetApiWithCredentials()

	tweet, err := api.PostTweet(text, nil)
	if err != nil {
		log.Print(err)
	}

	log.Print("the tweet posted! Contents:" + tweet.Text)
}
