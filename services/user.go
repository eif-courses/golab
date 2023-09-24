package services

import (
	"context"
	"time"
)

// https://www.sohamkamani.com/golang/jwt-authentication/
type User struct {
	UserId    string    `json:"user_id"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) GetAllUsers() ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT user_id, name, email, image, created_at, updated_at from users`

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	var users []*User
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.UserId,
			&user.Name,
			&user.Email,
			&user.Image,
			&user.CreatedAt,
			&user.UpdatedAt)

		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
func (u *User) CreateUser(user User) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `INSERT INTO users (name, email, password, image, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, $5, $6) returning *`

	_, err := db.ExecContext(ctx, query, user.Name, user.Email, user.Password, user.Image, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &user, nil

}

func (u *User) GetUserById(id string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `SELECT id, name, email, password, image, created_at, updated_at from users WHERE user_id=$1`

	var user User

	rows := db.QueryRowContext(ctx, query, id)
	err := rows.Scan(
		&user.UserId,
		&user.Name,
		&user.Image,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) UpdateUser(id string, body User) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `UPDATE users SET name=$1, email=$2, password=$3, image=$4, updated_at=$5 WHERE user_id=$6 returning *`

	_, err := db.ExecContext(ctx, query, body.Name, body.Email, body.Password, body.Image, time.Now(), id)
	if err != nil {
		return nil, err
	}
	return &body, nil
}

func (u *User) DeleteUser(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `DELETE FROM users WHERE user_id=$1`

	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
