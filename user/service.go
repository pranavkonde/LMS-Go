package user

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/pranavkonde/LMS-Go/db"
	"go.uber.org/zap"
)

type UserService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

type JWTClaim struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

var jwtKey = []byte("jsd549$^&")

type Service interface {
	List(ctx context.Context) (response ListResponse, err error)
	Create(ctx context.Context, req CreateRequest) (err error)
	FindByID(ctx context.Context, id string) (response FindByIDResponse, err error)
	DeleteByID(ctx context.Context, id string) (err error)
	Update(ctx context.Context, req UpdateRequest) (err error)
	GenerateJWT(ctx context.Context, Email string, Password string) (tokenString string, err error)
}

func (cs *UserService) GenerateJWT(ctx context.Context, Email string, Password string) (tokenString string, err error) {

	// var cs *userService
	user, err := cs.store.FindUserByEmail(ctx, Email)
	if err == db.ErrUserNotExist {
		cs.logger.Error("No user present", "err", err.Error())
		return "", errNoUserId
	}
	if err != nil {
		cs.logger.Error("Error finding user", "err", err.Error(), "email", Email)
		return
	}

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Id:    user.ID,
		Email: user.Email,
		Role:  user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func (cs *UserService) List(ctx context.Context) (response ListResponse, err error) {
	users, err := cs.store.ListUsers(ctx)
	if err == db.ErrUserNotExist {
		cs.logger.Error("No user present", "err", err.Error())
		return response, errNoUsers
	}
	if err != nil {
		cs.logger.Error("Error listing users", "err", err.Error())
		return
	}

	response.User = users
	return
}

func (cs *UserService) Create(ctx context.Context, c CreateRequest) (err error) {
	err = c.Validate()
	if err != nil {
		cs.logger.Errorw("Invalid request for user create", "msg", err.Error(), "user", c)
		return
	}
	uuidgen := uuid.New()
	c.ID = uuidgen.String()
	err = cs.store.CreateUser(ctx, &db.User{
		ID:        c.ID,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Gender:    c.Gender,
		Address:   c.Address,
		Age:       c.Age,
		Email:     c.Email,
		Password:  c.Password,
		MobileNum: c.MobileNum,
		Role:      c.Role,
	})
	if err != nil {
		cs.logger.Error("Error creating user", "err", err.Error())
		return
	}
	return
}

func (cs *UserService) Update(ctx context.Context, c UpdateRequest) (err error) {
	err = c.Validate()
	if err != nil {
		cs.logger.Error("Invalid Request for user update", "err", err.Error(), "user", c)
		return
	}

	err = cs.store.UpdateUser(ctx, &db.User{
		ID:        c.ID,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Gender:    c.Gender,
		Address:   c.Address,
		Age:       c.Age,
		Password:  c.Password,
		MobileNum: c.MobileNum,
	})
	if err != nil {
		cs.logger.Error("Error updating user", "err", err.Error(), "user", c)
		return
	}

	return
}

func (cs *UserService) FindByID(ctx context.Context, id string) (response FindByIDResponse, err error) {
	user, err := cs.store.FindUserByID(ctx, id)
	if err == db.ErrUserNotExist {
		cs.logger.Error("No user present", "err", err.Error())
		return response, errNoUserId
	}
	if err != nil {
		cs.logger.Error("Error finding user", "err", err.Error(), "id", id)
		return
	}

	response.User = user
	return
}

func (cs *UserService) DeleteByID(ctx context.Context, id string) (err error) {
	err = cs.store.DeleteUserByID(ctx, id)
	if err == db.ErrUserNotExist {
		cs.logger.Error("user Not present", "err", err.Error(), "id", id)
		return errNoUserId
	}
	if err != nil {
		cs.logger.Error("Error deleting user", "err", err.Error(), "id", id)
		return
	}

	return
}

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &UserService{
		store:  s,
		logger: l,
	}
}
