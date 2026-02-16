package interfaces

import "twitter-lld/internal/domain"

type FeedService interface {
	GetFeedStrategy(feedStrategy domain.FeedStrategy) ILoadFeedStrategy
	LoadTimeline(feed ILoadFeedStrategy)
}

