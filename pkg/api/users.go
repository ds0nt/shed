package api

import (
	"net/http"
	"time"

	"github.com/ds0nt/shed/domain/users"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// Securing the password using bcrypt
func (s *Service) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// Checking the password against the hashed password stored in the storage
func (s *Service) checkPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// Register - create new user
func (s *Service) createUserHandler(c echo.Context) error {
	user := &users.User{}
	if err := c.Bind(user); err != nil {
		return err
	}

	// hash the password and ignore the original one
	hashedPassword, err := s.hashPassword(user.Password)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Password encryption process failed")
	}
	user.Password = hashedPassword

	key := users.NewUserKey(user.Email)
	err = s.Store.CreateJSON(c.Request().Context(), "users", key.String(), &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// create new jwt token
	token, err := s.createJwtToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could not login")
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"token": token,
	})
}

// Login - get user info and return jwt token
func (s *Service) loginUserHandler(c echo.Context) error {
	user := &users.User{}
	if err := c.Bind(user); err != nil {
		return err
	}

	key := users.NewUserKey(user.Email)
	storedUser := users.User{}
	err := s.Store.ReadJSON(c.Request().Context(), "users", key.String(), &storedUser)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Username or password does not match")
	}
	if err := s.checkPassword(user.Password, storedUser.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, "Username or password does not match")
	}

	// create jwt token
	token, err := s.createJwtToken(&storedUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could not login")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

// userIdentity provides user related info in jwt token
type userIdentity struct {
	Username string `json:"usr"`
	UserID   int    `json:"uid"`
}

// secret helper variable, it is used to sign jwt token
var secret = []byte("my_secretfdsfsdfe44563t")

// createJwtToken creates a new JWT token for auth
func (s *Service) createJwtToken(user *users.User) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := t.SignedString(secret)
	if err != nil {
		return "", err
	}
	return token, nil
}
