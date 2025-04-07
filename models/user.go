package models

import (
	"database/sql"
	"errors"
	"go-jwt/database"
)

type User struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"-"`
}

func CreateUser(user *User) error {
	query := `INSERT INTO usuarios (nome, email, senha) VALUES ($1, $2, $3)`
	_, err := database.DB.Exec(query, user.Nome, user.Email, user.Senha)
	return err
}

func GetAllUsers() ([]User, error) {
	rows, err := database.DB.Query("SELECT id, nome, email FROM usuarios")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Nome, &u.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func GetUserByEmail(email string) (*User, error) {
	var u User
	row := database.DB.QueryRow("SELECT id, nome, email, senha FROM usuarios WHERE email = $1", email)
	err := row.Scan(&u.ID, &u.Nome, &u.Email, &u.Senha)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func GetUserByID(id int) (*User, error) {
	var u User
	row := database.DB.QueryRow("SELECT id, nome, email FROM usuarios WHERE id = $1", id)
	err := row.Scan(&u.ID, &u.Nome, &u.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &u, nil
}

func UpdateUser(id int, nome, email, senha string) error {
	query := `UPDATE usuarios SET nome=$1, email=$2, senha=$3 WHERE id=$4`
	_, err := database.DB.Exec(query, nome, email, senha, id)
	return err
}

func DeleteUser(id int) error {
	_, err := database.DB.Exec("DELETE FROM usuarios WHERE id = $1", id)
	return err
}
