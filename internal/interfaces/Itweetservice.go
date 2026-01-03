package interfaces

import "twitter-lld/internal/domain"

type TweetService interface { //Dependency Inversion principle :High-level modules should not depend on low-level modules. Both should depend on abstractions
	GetTweetMap() map[int]*domain.Tweet //Implemented interface segragation principle because we are just exposing the method that is needed.
}
