package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"

	keys "github.com/RyuseiNomi/shinkansen_bot/keys"
	text "github.com/RyuseiNomi/shinkansen_bot/text"
)

func Tweet() {
	text := text.GetTweetText()
	api := keys.GetApiWithCredentials()

	tweet, err := api.PostTweet(text, nil)
	if err != nil {
		log.Print(err)
	}

	log.Print("the tweet posted! Contents:" + tweet.Text)
}

func main() {
	lambda.Start(Tweet)
}
