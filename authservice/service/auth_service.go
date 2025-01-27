package service

import (
	"authservice/models"
	"authservice/repository"
	"context"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	AuthRepository *repository.AuthRepository
}

func NewAuthService(ar *repository.AuthRepository) *AuthService {
	return &AuthService{
		AuthRepository: ar,
	}
}

func (as *AuthService) Register(ctx context.Context, user *models.UserAuth) error {
	var wg sync.WaitGroup
	var err error
	wg.Add(1)

	go func() {
		defer wg.Done()
		err = as.AuthRepository.CreateNewUserAuth(ctx, user)
	}()
	wg.Wait()

	if err != nil {
		return err
	}
	return nil
}

func (as *AuthService) Login(ctx context.Context, indexType bool, index string, password string) (*models.User, string, error) {
	var wg sync.WaitGroup
	var err error
	var auth *models.Auth
	var user *models.User
	wg.Add(1)

	/* Retrieve auth information */
	if indexType {
		go func() {
			defer wg.Done()
			auth, err = as.AuthRepository.GetAuthByEmail(ctx, index)
		}()
	} else {
		go func() {
			defer wg.Done()
			auth, err = as.AuthRepository.GetAuthByUsername(ctx, index)
		}()
	}

	/* Hash input password and compare it to stored password */
	hashedPassword, err := hashPassword(password, auth.Salt)
	if err != nil {
		return nil, "", err
	} else if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(auth.Password)); err != nil {
		return nil, "", err
	}

	/* Create JWT signed string */
	tokenString, err := generateJWT(auth.ID)
	if err != nil {
		return nil, "", err
	}

	/* Query database for relevant user details */
	wg.Add(1)
	go func() {
		defer wg.Done()
		user, err = as.AuthRepository.GetUserDetails(ctx, auth.ID)
	}()
	wg.Wait()

	if err != nil {
		return nil, "", err
	}
	return user, tokenString, nil
}

func (as *AuthService) Logout(ctx context.Context, id uint64) error { return nil }

func (as *AuthService) UpdateUsername(ctx context.Context, id uint64, username string) error {
	var wg sync.WaitGroup
	var err error
	wg.Add(1)

	go func() {
		defer wg.Done()
		err = as.UpdateUsername(ctx, id, username)
	}()
	wg.Wait()

	if err != nil {
		return err
	}
	return nil
}

func (as *AuthService) UpdatePassword(ctx context.Context, id uint64, password string) error {
	var wg sync.WaitGroup
	var err error
	wg.Add(1)

	go func() {
		defer wg.Done()
		err = as.UpdateUsername(ctx, id, password)
	}()
	wg.Wait()

	if err != nil {
		return err
	}
	return nil
}

func (as *AuthService) DeleteUser(ctx context.Context, id uint64) error {
	var wg sync.WaitGroup
	var err error
	wg.Add(1)

	go func() {
		defer wg.Done()
		err = as.DeleteUser(ctx, id)
	}()
	wg.Wait()

	if err != nil {
		return err
	}
	return nil
}
