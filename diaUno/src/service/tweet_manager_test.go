package service_test

import (
	"github.com/faramirez18/GO/diaUno/src/service"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	tweet := "Este es mi primer tweet"

	service.PublishTweet(tweet)

	if service.GetTweet() != tweet {
		t.Error("Expected tweet is", tweet)
	}
}