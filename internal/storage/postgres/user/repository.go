package user

import (
	"context"
	"errors"
	"time"

	"github.com/aygumov-g/service-users-go/internal/domain/user"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

func (r *Repository) GetByID(ctx context.Context, id int64) (*user.User, error) {
	const query = `
		SELECT
			id,
			name,
			avatar_url,
			bio,
			created_at,
			updated_at
		FROM users
		WHERE id = $1
	`

	row := r.pool.QueryRow(ctx, query, id)

	var u user.User
	err := row.Scan(
		&u.ID,
		&u.Name,
		&u.AvatarURL,
		&u.Bio,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *Repository) Create(ctx context.Context, u *user.User) error {
	const query = `
		INSERT INTO users (
			id,
			name,
			avatar_url,
			bio,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.pool.Exec(
		ctx,
		query,
		u.ID,
		u.Name,
		u.AvatarURL,
		u.Bio,
		time.Now().UTC(),
		time.Now().UTC(),
	)

	return err
}

func (r *Repository) Update(ctx context.Context, u *user.User) error {
	const query = `
		UPDATE users
		SET
			name = $2,
			avatar_url = $3,
			bio = $4,
			updated_at = $5
		WHERE id = $1
	`

	_, err := r.pool.Exec(
		ctx,
		query,
		u.ID,
		u.Name,
		u.AvatarURL,
		u.Bio,
		time.Now().UTC(),
	)

	return err
}
