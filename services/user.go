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

	query := `INSERT INTO users (name, image, created_at, updated_at)
			  VALUES ($1, $2, $3, $4) returning *`

	_, err := db.ExecContext(ctx, query, user.Name, user.Image, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &user, nil

}

func (u *User) GetUserById(id string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `SELECT id, name, image, created_at, updated_at from users WHERE id=$1`

	var user User

	rows := db.QueryRowContext(ctx, query, id)
	err := rows.Scan(
		&user.ID,
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

	query := `UPDATE users SET name=$1, image=$2, updated_at=$3 WHERE id=$4 returning *`

	_, err := db.ExecContext(ctx, query, body.Name, body.Image, time.Now(), id)
	if err != nil {
		return nil, err
	}
	return &body, nil
}

func (u *User) DeleteUser(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `DELETE FROM users WHERE id=$1`

	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
