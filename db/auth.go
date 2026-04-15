package db

import (
	"errors"
	"taskflow-samrat/models"
	"taskflow-samrat/resources"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(payload models.UserRegister) (models.User, error) {
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id, name, email, created_at`
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), 12)
	row := resources.DB.QueryRow(query, payload.Name, payload.Email, hashedPassword)
	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func LoginUser(payload models.UserLogin) (bool, models.User, error) {
	query := `SELECT id,name,email,password,created_at FROM users WHERE email = $1`
	row := resources.DB.QueryRow(query, payload.Email)
	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return false, user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return false, user, errors.New("invalid credentials")
	}
	return true, user, nil
}
