package domain

import (
	"fmt"
	"time"
)

type Tweet struct {
	Id       int
	Content  string
	UserId   int
	UserName string
	Time     time.Time
	Likes    []int
	Comments []Comment
}

func (t *Tweet) AddCommment(comment Comment) {
	t.Comments = append(t.Comments, comment)
	fmt.Printf("%v added a comment : %v on the tweet '%v' by %v\n", comment.UserName, comment.Content, t.Content, t.UserName)
}

func (t *Tweet) Like(Userid int, name string) {
	t.Likes = append(t.Likes, Userid)
	fmt.Printf("%v liked the tweet '%v' by %v\n", name, t.Content, t.UserName)
}

var (
	TweetId = 0
)

func GetTweetId() int {
	TweetId++
	return TweetId
}
