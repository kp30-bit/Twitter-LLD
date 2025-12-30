package usecase

import (
	"sort"
	"twitter-lld/internal/domain"
)

type PopularitySortedFeed struct {
	Tweets map[int]*domain.Tweet
}

func (t PopularitySortedFeed) LoadFeed() []*domain.Tweet {
	tweetSlice := make([]*domain.Tweet, 0, len(t.Tweets))

	for _, t := range t.Tweets {
		tweetSlice = append(tweetSlice, t)
	}

	sort.Slice(tweetSlice, func(i, j int) bool {
		return len(tweetSlice[i].Likes) > len(tweetSlice[j].Likes)
	})

	return tweetSlice
}
