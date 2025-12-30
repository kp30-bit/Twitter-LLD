package usecase

import (
	"sort"
	"twitter-lld/internal/domain"
)

type TimeSortedFeed struct {
	Tweets map[int]*domain.Tweet
}

func (t *TimeSortedFeed) LoadFeed() []*domain.Tweet {
	tweetSlice := make([]*domain.Tweet, 0, len(t.Tweets))

	for _, t := range t.Tweets {
		tweetSlice = append(tweetSlice, t)
	}

	sort.Slice(tweetSlice, func(i, j int) bool {
		return tweetSlice[i].Time.After(tweetSlice[j].Time)
	})

	return tweetSlice
}
