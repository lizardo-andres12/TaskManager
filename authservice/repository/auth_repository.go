package repository

import (
	"context"
	"database/sql"

	"authservice/models"
)

type AuthRepository struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		DB: db,
	}
}

func (ar *AuthRepository) CreateNewUserAuth(ctx context.Context, auth *models.UserAuth) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, err := ar.DB.ExecContext(
		ctx,
		"INSERT INTO auth (username, email, password, salt, role, firstName, lastName, mediaURL) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		auth.Username,
		auth.Email,
		auth.Password,
		auth.Salt,
		auth.Role,
		auth.FirstName,
		auth.LastName,
		auth.MediaURL,
	)
	if err != nil {
		return err
	}
	return nil
}

func (ar *AuthRepository) GetUserDetails(ctx context.Context, id uint64) (*models.User, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	var user models.User
	row := ar.DB.QueryRowContext(
		ctx,
		"SELECT role, firstName, lastName, mediaUrl FROM auth WHERE id = ?",
		id,
	)

	if err := row.Scan(&user.Role, &user.FirstName, &user.LastName, &user.MediaURL); err != nil {
		return nil, err
	}
	return &user, nil
}

func (ar *AuthRepository) GetAuthByEmail(ctx context.Context, email string) (*models.Auth, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	var auth models.Auth
	row := ar.DB.QueryRowContext(ctx, "SELECT id, password, salt FROM auth WHERE email = ?", email)

	if err := row.Scan(&auth.ID, &auth.Password, &auth.Salt); err != nil {
		return nil, err
	}
	return &auth, nil
}

func (ar *AuthRepository) GetAuthByUsername(ctx context.Context, username string) (*models.Auth, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	var auth models.Auth
	row := ar.DB.QueryRowContext(ctx, "SELECT id, password, salt FROM auth WHERE username = ?", username)

	if err := row.Scan(&auth.ID, &auth.Password, &auth.Salt); err != nil {
		return nil, err
	}
	return &auth, nil
}

func (ar *AuthRepository) UpdateUsername(ctx context.Context, id uint64, username string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, err := ar.DB.ExecContext(ctx, "UPDATE auth SET username = ? WHERE id = ?", username, id)
	if err != nil {
		return err
	}
	return nil
}

func (ar *AuthRepository) UpdatePassword(ctx context.Context, id uint64, password string, salt string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, err := ar.DB.ExecContext(ctx, "UPDATE auth SET password = ?, salt = ? WHERE id = ?", password, salt, id)
	if err != nil {
		return err
	}
	return nil
}

func (ar *AuthRepository) DeleteUser(ctx context.Context, id uint64) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, err := ar.DB.ExecContext(ctx, "DELETE FROM auth WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
