package main

import (
	"time"
	"twitter-lld/internal/domain"
	"twitter-lld/internal/services"
)

func main() {
	Alice := &domain.User{
		Id:        domain.GetUserId(),
		Name:      "Alice",
		Followers: make(map[int]*domain.User),
	}
	Bob := &domain.User{
		Id:        domain.GetUserId(),
		Name:      "Bob",
		Followers: make(map[int]*domain.User),
	}

	// construct concrete services
	tweetService := services.NewTweetService()
	userService := services.NewUserService()
	feedService := services.NewFeedService(tweetService)

	// inject them into the Twitter facade via interfaces
	Twitter := services.NewTwitter(tweetService, userService, feedService)
	Twitter.UserService.AddUser(Alice)
	Twitter.UserService.AddUser(Bob)
	Twitter.Follow(Alice.Id, Bob.Id)
	tweetId := Twitter.Tweet(domain.Tweet{
		Id:       domain.GetTweetId(),
		Content:  "Hi Followers!!!",
		UserId:   Alice.Id,
		UserName: Alice.Name,
		Time: time.Date(
			2025, time.January, 30, // Year, Month, Day
			15, 04, 05, // Hour, Min, Sec
			0,          // Nanoseconds
			time.Local, // or time.UTC
		),
	})
	Twitter.Like(Bob.Id, Bob.Name, tweetId)
	Twitter.Comment(tweetId, &domain.Comment{
		Id:       domain.GetCommentId(),
		Content:  "Hi Alice :) \n",
		UserId:   Bob.Id,
		UserName: Bob.Name,
		TweetId:  tweetId,
	})
	Twitter.Tweet(domain.Tweet{
		Id:       domain.GetTweetId(),
		Content:  "First tweet!\n",
		UserId:   Bob.Id,
		UserName: Bob.Name,
		Time: time.Date(
			2027, time.January, 30, // Year, Month, Day
			15, 04, 05, // Hour, Min, Sec
			0,          // Nanoseconds
			time.Local, // or time.UTC
		),
	})

	Twitter.LoadTimeline(domain.TimeSortedFeed)
}
