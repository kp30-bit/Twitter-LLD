package services

import (
	"sync"
	"twitter-lld/internal/domain"
)

type Twitter struct {
	FeedService  *FeedService
	TweetService *TweetService
	UserService  *UserService
}

var (
	TwitterInst *Twitter
	TwitterOnce sync.Once
)

func NewTwiter() *Twitter {
	TwitterOnce.Do(func() {
		TwitterInst = &Twitter{
			TweetService: NewTweetService(),
			UserService:  NewUserService(),
		}
		TwitterInst.FeedService = NewFeedService(TwitterInst.TweetService)
	})
	return TwitterInst
}

func (t *Twitter) Tweet(tweet domain.Tweet) int {
	return t.TweetService.AddTweet(&tweet)
}

func (t *Twitter) Comment(tweetId int, comment *domain.Comment) {
	t.TweetService.AddComment(tweetId, comment)
}

func (t *Twitter) Like(userId int, name string, tweetId int) {
	t.TweetService.Like(tweetId, name, tweetId)
}

func (t *Twitter) Follow(id1, id2 int) {
	t.UserService.Follow(id1, id2)

}

func (t *Twitter) UnFollow(id1, id2 int) {
	t.UserService.UnFollow(id1, id2)

}

func (t *Twitter) LoadTimeline(strategy domain.FeedStrategy) {
	feedstrategy := t.FeedService.GetFeedStrategy(strategy)
	t.FeedService.LoadTimeline(feedstrategy)
}
