package repository

import (
	"database/sql"
	"golang-backend/internal/model"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) Create(user model.User) error {
	query := `INSERT INTO users(name,email,password) VALUES($1,$2,$3)`
	_, err := r.DB.Exec(query, user.Name, user.Email, user.Password)
	return err
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	row := r.DB.QueryRow("SELECT id,name,email,password FROM users WHERE email=$1", email)

	var user model.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, nil
	}

	return &user, nil
}
