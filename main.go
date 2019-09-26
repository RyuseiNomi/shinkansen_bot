package main

import (
	"log"

	keys "./keys"
	text "./text"
)

func main() {

	text := text.GetTweetText()
	api := keys.GetApiWithCredentials()

	tweet, err := api.PostTweet(text, nil)
	if err != nil {
		log.Print(err)
	}

	log.Print("the tweet posted! Contents:" + tweet.Text)
}
