package main

import "time"

type Tweet struct {
	user string
	text string
	Date *time.Time
}

func NewTweet(user , text string) *Tweet {

	date := time.Now()

	tweet := Tweet{
		user,
		text,
		&date,
	}

	return &tweet
}


