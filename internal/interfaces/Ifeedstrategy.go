package interfaces

import "twitter-lld/internal/domain"

type ILoadFeedStrategy interface {
	LoadFeed() []*domain.Tweet
}
