package interfaces

import "twitter-lld/internal/domain"

type UserService interface {
	AddUser(user *domain.User)
	Follow(followeeId, followerId int) error
	UnFollow(followeeId, followerId int) error
	GetAllFollowers(userId int)
}

