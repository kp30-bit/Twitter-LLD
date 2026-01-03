package interfaces

import "twitter-lld/internal/domain"

type ILoadFeedStrategy interface { //Dependency Infversion principle :High-level modules should not depend on low-level modules. Both should depend on abstractions
	LoadFeed() []*domain.Tweet //Implemented interface segragation principle because we are just exposing the method that is needed.
}
