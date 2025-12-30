package domain

type Comment struct {
	Id       int
	Content  string
	UserId   int
	TweetId  int
	UserName string
}

var (
	CommentId = 0
)

func GetCommentId() int {
	CommentId++
	return CommentId
}
