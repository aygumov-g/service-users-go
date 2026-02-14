package user

import (
	"context"
	"errors"

	d_user "github.com/aygumov-g/service-users-go/internal/domain/user"
	srv_user "github.com/aygumov-g/service-users-go/internal/service/user"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

// func (r *Repository) Upsert(ctx context.Context, userIn *user.User) (*user.User, error) {
// 	row := r.db.QueryRow(
// 		ctx,
// 		`
// 		INSERT INTO users (
// 			id,
// 			first_name,
// 			last_name,
// 			bio,
// 			avatar_url,
// 			created_at,
// 			updated_at
// 		)
// 		VALUES ($1, $2, $3, $4, $5, $6, $7)
// 		ON CONFLICT (id) DO UPDATE
// 		SET
// 			first_name = EXCLUDED.first_name,
// 			last_name = EXCLUDED.last_name,
// 			bio = EXCLUDED.bio,
// 			avatar_url = EXCLUDED.avatar_url,
// 			updated_at = EXCLUDED.updated_at
// 		RETURNING
// 			id,
// 			first_name,
// 			last_name,
// 			bio,
// 			avatar_url,
// 			created_at,
// 			updated_at
// 		`,
// 		userIn.ID,
// 		userIn.FirstName,
// 		userIn.LastName,
// 		userIn.Bio,
// 		userIn.AvatarURL,
// 		userIn.CreatedAt,
// 		userIn.UpdatedAt,
// 	)

// 	var userOut user.User
// 	if err := row.Scan(
// 		&userOut.ID,
// 		&userOut.FirstName,
// 		&userOut.LastName,
// 		&userOut.Bio,
// 		&userOut.AvatarURL,
// 		&userOut.CreatedAt,
// 		&userOut.UpdatedAt,
// 	); err != nil {
// 		return nil, err
// 	}

// 	return &userOut, nil
// }

func (r *Repository) GetByID(ctx context.Context, id int64) (*d_user.User, error) {
	row := r.db.QueryRow(
		ctx,
		`
		SELECT
			id,
			first_name,
			last_name,
			bio,
			avatar_url,
			created_at,
			updated_at
		FROM users
		WHERE id = $1
		`,
		id,
	)

	var user d_user.User
	if err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Bio,
		&user.AvatarURL,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, srv_user.ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (r *Repository) Create(ctx context.Context, user *d_user.User) error {
	_, err := r.db.Exec(
		ctx,
		`
		INSERT INTO users (
			id,
			first_name,
			last_name,
			bio,
			avatar_url,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		`,
		user.ID,
		user.FirstName,
		user.LastName,
		user.Bio,
		user.AvatarURL,
		user.CreatedAt,
		user.UpdatedAt,
	)

	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				return srv_user.ErrUserAlreadyExists
			}
		}

		return err
	}

	return nil
}

func (r *Repository) Update(ctx context.Context, user *d_user.User) error {
	_, err := r.db.Exec(
		ctx,
		`
		UPDATE users
		SET
			first_name = $2,
			last_name = $3,
			bio = $4,
			avatar_url = $5,
			updated_at = $6
		WHERE id = $1
		`,
		user.ID,
		user.FirstName,
		user.LastName,
		user.Bio,
		user.AvatarURL,
		user.UpdatedAt,
	)

	return err
}
