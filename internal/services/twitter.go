package services

import (
	"twitter-lld/internal/domain"
	"twitter-lld/internal/interfaces"
)

type Twitter struct {
	FeedService  interfaces.FeedService
	TweetService interfaces.TweetService
	UserService  interfaces.UserService
}

// NewTwitter constructs a Twitter facade from its dependencies.
// Using interface types here respects the Dependency Inversion Principle.
func NewTwitter(
	tweetService interfaces.TweetService,
	userService interfaces.UserService,
	feedService interfaces.FeedService,
) *Twitter {
	return &Twitter{
		TweetService: tweetService,
		UserService:  userService,
		FeedService:  feedService,
	}
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
