package domain

import "fmt"

type User struct {
	Id        int
	Name      string
	Followers map[int]*User
}

var (
	UserId = 0
)

func GetUserId() int {
	UserId++
	return UserId
}

func (u *User) AddFollower(Follower *User) {
	u.Followers[Follower.Id] = Follower
	fmt.Printf("%v follows %v now\n", u.Name, Follower.Name)
}

func (u *User) RemoveFollower(Follower *User) {
	delete(u.Followers, Follower.Id)
}

func (u *User) GetAllFollowers() {
	fmt.Printf("List of Followers of %v\n", u.Name)

	for _, v := range u.Followers {
		fmt.Printf("%v | ", v.Name)
	}
	fmt.Printf("\n")
}
