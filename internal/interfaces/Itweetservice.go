package interfaces

import "twitter-lld/internal/domain"

type TweetService interface {
	GetTweetMap() map[int]*domain.Tweet
}
