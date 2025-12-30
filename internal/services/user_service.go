package services

import (
	"fmt"
	"twitter-lld/internal/domain"
)

type UserService struct {
	UserList map[int]*domain.User
}

func NewUserService() *UserService {
	return &UserService{
		UserList: make(map[int]*domain.User),
	}
}

func (us *UserService) AddUser(user *domain.User) {
	us.UserList[user.Id] = user
}

func (us *UserService) Follow(followeeId, followerId int) error {
	follower, ok := us.UserList[followerId]
	if !ok {
		return fmt.Errorf("follower does not exist, could not follow\n")
	}
	followee, ok := us.UserList[followeeId]
	if !ok {
		return fmt.Errorf("followee does not exist, could not follow\n")
	}
	follower.AddFollower(followee)
	return nil
}

func (us *UserService) UnFollow(followeeId, followerId int) error {
	follower, ok := us.UserList[followeeId]
	if !ok {
		return fmt.Errorf("follower does not exist, could not unfollow\n")
	}
	followee, ok := us.UserList[followeeId]
	if !ok {
		return fmt.Errorf("followee does not exist, could not unfollow\n")
	}
	follower.RemoveFollower(followee)
	return nil
}

func (us *UserService) GetAllFollowers(userId int) {
	user := us.UserList[userId]

	user.GetAllFollowers()
}
