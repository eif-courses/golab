package services

import (
	"context"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) GetAllUsers() ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, name, image, created_at, updated_at from users`

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	var users []*User
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.Name,
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

	query := `INSERT INTO users (id, name, image, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, $5) return *`

	_, err := db.ExecContext(ctx, query, user.ID, user.Name, user.Image, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &user, nil

}
