package user

import "time"

type User struct {
	ID        string
	Name      string
	AvatarURL string
	Bio       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
