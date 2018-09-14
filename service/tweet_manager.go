package service

import "testing"

var Tweet string

func PublishTweet(tweet *domain.Tweet) {
	Tweet = tweet
}

func GetTweet() *domain.Tweet {
	return Tweet
}

func TestPublishedTweetIsSaved(t *testing.T) {

	var tweet string = "This is my first tweet"

	service.PublishTweet(tweet)

	if service.Tweet != tweet {
		t.Error("Expected tweet is", tweet)
	}

}