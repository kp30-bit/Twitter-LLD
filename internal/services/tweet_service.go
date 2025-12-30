package services

import (
	"fmt"
	"twitter-lld/internal/domain"
)

type TweetService struct {
	Tweets map[int]*domain.Tweet
}

func NewTweetService() *TweetService {
	return &TweetService{
		Tweets: make(map[int]*domain.Tweet),
	}
}

func (ts *TweetService) AddTweet(tweet *domain.Tweet) int {
	ts.Tweets[tweet.Id] = tweet
	fmt.Printf("%v tweeted : %v\n", tweet.UserName, tweet.Content)
	return tweet.Id
}

func (tc *TweetService) Like(userId int, name string, tweetId int) error {
	tweet, ok := tc.Tweets[tweetId]
	if !ok {
		return fmt.Errorf("Could not like tweet : %v as it is not available anymore\n", tweetId)
	}
	tweet.Like(userId, name)
	return nil
}

func (tc *TweetService) AddComment(tweetId int, comment *domain.Comment) error {
	tweet, ok := tc.Tweets[tweetId]
	if !ok {
		return fmt.Errorf("Could not like tweet : %v as it is not available anymore\n", tweetId)
	}
	tweet.AddCommment(*comment)
	return nil
}

func (tc *TweetService) GetTweetMap() map[int]*domain.Tweet {
	copyMap := make(map[int]*domain.Tweet)

	for _, v := range tc.Tweets {
		copyMap[v.Id] = v
	}
	return copyMap
}
