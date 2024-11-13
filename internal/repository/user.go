package repository

import (
	"ServiceT/internal/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Импорт драйвера PostgreSQL
)

// UserRepository представляет доступ к данным пользователей
type UserRepository struct {
	db *sqlx.DB
}

// NewUserRepository создает новый экземпляр UserRepository
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create добавляет нового пользователя в базу данных
func (r *UserRepository) Create(user *model.User) error {
	_, err := r.db.NamedExec(`INSERT INTO users (username, password) VALUES (:username, :password)`, user)
	return err
}

// FindByUsernameAndPassword находит пользователя по имени пользователя и паролю
func (r *UserRepository) FindByUsernameAndPassword(username, password string) (*model.User, error) {
	var user model.User
	err := r.db.Get(&user, `SELECT id, username FROM users WHERE username=$1 AND password=$2`, username, password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByID находит пользователя по ID
func (r *UserRepository) FindByID(id string) (*model.User, error) {
	var user model.User
	err := r.db.Get(&user, `SELECT id, username FROM users WHERE id=$1`, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
