package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/erjiridholubis/go-superindo-product/internal/model"
)

var (
	// Query Get User By Username
	QueryGetUserByUsername = `SELECT id, name, username, password FROM users WHERE username = $1`
)

// GetUserByUsername is a function to get user data by username from database
func (pr *postgreRepository) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User

	err := pr.ConnDB.QueryRowContext(ctx, QueryGetUserByUsername, username).Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.Password,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}