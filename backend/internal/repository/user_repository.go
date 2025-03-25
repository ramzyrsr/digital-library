package repository

import (
	"database/sql"

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
