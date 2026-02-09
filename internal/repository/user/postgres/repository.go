package postgres

import (
	"context"
	"errors"

	"github.com/aygumov-g/service-users-go/internal/domain/user"
	u "github.com/aygumov-g/service-users-go/internal/service/user"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetByID(ctx context.Context, id int64) (*user.User, error) {
	row := r.db.QueryRow(
		context.Background(),
		`
		SELECT
			id,
			login,
			created_at,
			updated_at
		FROM users
		WHERE id = $1
		`,
		id,
	)

	var user user.User
	if err := row.Scan(
		&user.ID,
		&user.Login,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, u.ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (r *Repository) Create(ctx context.Context, user *user.User) error {
	_, err := r.db.Exec(
		context.Background(),
		`
		INSERT INTO users (
			id,
			login,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4)
		`,
		user.ID,
		user.Login,
		user.CreatedAt,
		user.UpdatedAt,
	)

	return err
}

func (r *Repository) Update(ctx context.Context, user *user.User) error {
	_, err := r.db.Exec(
		context.Background(),
		`
		UPDATE users
		SET
			login = $1,
			updated_at = $2
		WHERE id = $3
		`,
		user.Login,
		user.UpdatedAt,
		user.ID,
	)

	return err
}
