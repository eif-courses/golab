package services

import (
	"context"
	"github.com/eif-courses/golab/utils"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte("my_secret_key")

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

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

func (u *User) SignIn(credentials Credentials) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT password FROM users WHERE email=$1`

	var user User

	rows := db.QueryRowContext(ctx, query, credentials.Username)
	err := rows.Scan(
		&user.Password,
	)

	if err != nil {
		return "", err
	}

	if utils.CheckPasswordHash(credentials.Password, user.Password) == false {
		return "", nil
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: credentials.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
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

	hashed, er := utils.GeneratehashPassword(user.Password)
	if er != nil {
		return nil, er
	}

	_, err := db.ExecContext(ctx, query, user.Name, user.Email, hashed, user.Image, user.CreatedAt, user.UpdatedAt)

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
