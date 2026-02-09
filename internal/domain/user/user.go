package user

import "time"

type User struct {
	ID        int64
	Login     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
