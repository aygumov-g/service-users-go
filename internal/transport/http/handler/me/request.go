package me

import "github.com/aygumov-g/service-users-go/internal/service/user"

type updateRequest struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Bio       *string `json:"bio"`
	AvatarURL *string `json:"avatar_url"`
}

func (r updateRequest) toServiceInput() user.UpdateInput {
	return user.UpdateInput{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Bio:       r.Bio,
		AvatarURL: r.AvatarURL,
	}
}
