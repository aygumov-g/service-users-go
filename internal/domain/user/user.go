package user

import "time"

type User struct {
	ID        int64
	FirstName string
	LastName  string
	Bio       string
	AvatarURL string
	CreatedAt time.Time
	UpdatedAt time.Time
}
