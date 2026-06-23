package service

import (
	"e-backend-boilerplate/modules/auth/models"
	"e-backend-boilerplate/modules/auth/repository"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var ErrAuthUserNotFound = errors.New("user is not found or password is incorrect")

type Service struct {
	repo         *repository.Repository
	jwtSecretKey string
}

func NewService(repo *repository.Repository, jwtSecretKey string) *Service {
	return &Service{repo, jwtSecretKey}
}

func (s *Service) Create(item models.User) (user *models.User, token string, err error) {
	// Hash password
	passwordHash, err := s.hashPassword(item.Password)
	if err != nil {
		return
	}
	item.Password = passwordHash

	user, err = s.repo.Create(item)
	if err != nil {
		return
	}

	token, err = s.generateUserJWTToken(*user)
	if err != nil {
		return
	}

	return
}

func (s *Service) Update(id uint, item models.User) (*models.User, error) {
	return s.repo.Update(id, item)
}

func (s *Service) Get(id uint) (*models.User, error) {
	return s.repo.Get(id)
}

func (s *Service) Delete(id uint) (err error) {
	return s.repo.Delete(id)
}

func (s *Service) SignIn(credential, password string) (user *models.User, token string, err error) {
	user, err = s.repo.FindByUsernameOrEmail(credential)
	if err != nil {
		err = ErrAuthUserNotFound
		return
	}

	if s.verifyPassword(password, user.Password) {
		token, err = s.generateUserJWTToken(*user)
		return
	}

	err = ErrAuthUserNotFound
	return
}

func (s *Service) hashPassword(password string) (hash string, err error) {
	hashBytes, err :=
		bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hash = string(hashBytes)
	return
}

func (s *Service) verifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *Service) generateUserJWTToken(user models.User) (token string, err error) {
	claims := &models.AuthClaims{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: strconv.FormatUint(uint64(user.ID), 10),
			// TODO: move ExpiresAt value to the config
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 48)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = t.SignedString([]byte(s.jwtSecretKey))
	return
}
