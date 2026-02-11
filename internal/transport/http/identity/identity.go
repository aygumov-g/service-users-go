package identity

import (
	"context"

	"github.com/aygumov-g/service-users-go/internal/domain/identity"
)

type Identity struct {
	contextTokenKey string
}

func NewIdentity(key string) *Identity {
	return &Identity{
		contextTokenKey: key,
	}
}

func (i *Identity) Upload(ctx context.Context, idntt *identity.Identity) context.Context {
	return context.WithValue(ctx, i.contextTokenKey, idntt)
}

func (i *Identity) Unload(ctx context.Context) (identity.Identity, bool) {
	idntt, ok := ctx.Value(i.contextTokenKey).(identity.Identity)
	return idntt, ok
}
