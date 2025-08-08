package domain

import "context"

type User struct {
	ID       int
	Name     string // Имя пользователя
	Email    string // Email пользователя
	Password string // Захешированный пароль
}

type UserRepository interface {
	// Создание нового пользователя в базе
	Create(ctx context.Context, user *User) error
	// Получение пользователя из базы по его ID
	GetByID(ctx context.Context, id int) (*User, error)
	// Получение пользователя из базы по его email
	GetByEmail(ctx context.Context, email string) (*User, error)
	// Изменение данных пользователя в базе данных
	Update(ctx context.Context, user *User) error
	// Удаление пользователя из базы
	Delete(ctx context.Context, id int) error
}
