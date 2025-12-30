package services

import (
	"fmt"
	"twitter-lld/internal/domain"
	"twitter-lld/internal/interfaces"
	"twitter-lld/internal/usecase"
)

type FeedService struct {
	TweetService interfaces.TweetService
}

func NewFeedService(TweetService interfaces.TweetService) *FeedService {
	return &FeedService{
		TweetService: TweetService,
	}
}

func (fs *FeedService) GetFeedStrategy(feedStrategy domain.FeedStrategy) interfaces.ILoadFeedStrategy {
	switch feedStrategy {
	case domain.PopularitySortedFeed:
		return usecase.PopularitySortedFeed{
			Tweets: fs.TweetService.GetTweetMap(),
		}
	case domain.TimeSortedFeed:
		return &usecase.TimeSortedFeed{
			Tweets: fs.TweetService.GetTweetMap(),
		}
	}
	return nil
}

func (fs *FeedService) LoadTimeline(feed interfaces.ILoadFeedStrategy) {
	feedTweets := feed.LoadFeed()
	fmt.Printf("------------- Timeline ----------- \n")
	for ind, tweet := range feedTweets {
		fmt.Printf("Tweet %v : %v\n", ind, tweet.Content)
	}
}
