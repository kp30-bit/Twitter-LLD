package domain

type FeedStrategy int

const (
	PopularitySortedFeed FeedStrategy = iota
	TimeSortedFeed
)
