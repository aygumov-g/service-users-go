package users

import "github.com/aygumov-g/service-users-go/internal/domain/user"

type userResponse struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Bio       string `json:"bio"`
	AvatarURL string `json:"avatar_url"`
}

func (r userResponse) toResponse(u *user.User) userResponse {
	return userResponse{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Bio:       u.Bio,
		AvatarURL: u.AvatarURL,
	}
}
