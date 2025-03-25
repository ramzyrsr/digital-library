package repository

import (
	"database/sql"
	"errors"

	"github.com/ramzyrsr/digital-library/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) CreateUser(user *entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (email, password, name, role) 
		VALUES
			($1, $2, $3, $4)`
	_, err = r.DB.Exec(query, user.Email, string(hashedPassword), user.Name, user.Role)

	return err
}

func (r *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	query := `SELECT id, email, password, name, role, created_at FROM users WHERE email = $1`
	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.Role, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}
