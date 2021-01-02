package domain

import (
	"sort"
	"time"
)

type user struct {
	id         int
	name       string
	email      string
	password   string
	bio        *string
	imageLink  *string
	followings []int
	favorites  []article
	createdAt  time.Time
	updatedAt  time.Time
}

type userUpdatableProperty int

const (
	UserEmail userUpdatableProperty = iota
	UserName
	UserBio
	UserImageLink
	UserPassword
)

func UpdateUser(initial *user, opts ...func(fields *user)) {
	for _, v := range opts {
		v(initial)
	}
}

func SetUserName(input *string) func(fields *user) {
	return func(initial *user) {
		if input != nil {
			initial.name = *input
		}
	}
}

func SetUserEmail(input *string) func(fields *user) {
	return func(initial *user) {
		if input != nil {
			initial.email = *input
		}
	}
}

// give empty string to delete it
func SetUserBio(input *string) func(fields *user) {
	return func(initial *user) {
		if input != nil {
			if *input == "" {
				initial.bio = nil
				return
			}
			initial.bio = input
		}
	}
}

// give empty string to delete it
func SetUserImageLink(input *string) func(fields *user) {
	return func(initial *user) {
		if input != nil {
			if *input == "" {
				initial.imageLink = nil
				return
			}
			initial.imageLink = input
		}
	}
}

func SetUserPassword(input *string) func(fields *user) {
	return func(initial *user) {
		if input != nil {
			initial.password = *input
		}
	}
}

func (user user) Follows(userName int) bool {
	if user.followings == nil {
		return false
	}

	sort.Ints(user.followings)
	i := sort.SearchInts(user.followings, userName)
	return i < len(user.followings) && user.followings[i] == userName
}

// UpdateFollowees will append or remove followee to current user according to follow param
func (user *user) UpdateFollowees(followeeName int, follow bool) {
	if follow {
		user.followings = append(user.followings, followeeName)
		return
	}

	for i := 0; i < len(user.followings); i++ {
		if user.followings[i] == followeeName {
			user.followings = append(user.followings[:i], user.followings[i+1:]...) // memory leak ? https://github.com/golang/go/wiki/SliceTricks
		}
	}
	if len(user.followings) == 0 {
		user.followings = nil
	}
}
