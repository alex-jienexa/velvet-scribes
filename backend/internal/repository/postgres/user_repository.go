package postgres

import (
	"context"

	"github.com/alex-jienexa/velvet-scribes/internal/domain"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	return r.db.QueryRow(ctx, query, user.Name, user.Email, user.Password).Scan(&user.ID)
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	query := `SELECT id, name, email, password FROM users WHERE id = $1`
	var user domain.User
	err := r.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err == pgx.ErrNoRows {
		return nil, domain.ErrUserNotFound // Доменная ошибка
	}
	return &user, err
}
