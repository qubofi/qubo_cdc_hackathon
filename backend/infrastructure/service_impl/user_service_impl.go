package service_impl

import (
	"database/sql"
	"errors"
	"qubo/qubo-backend/domain/entity"
	"qubo/qubo-backend/domain/service"
	"qubo/qubo-backend/infrastructure/database"
	"qubo/qubo-backend/infrastructure/database/dal"
)

type UserServiceImpl struct {
	db *sql.DB
}

func NewUserService() *UserServiceImpl {
	return &UserServiceImpl{
		db: database.Db,
	}
}

var _ service.UserService = &UserServiceImpl{}

func (userRepo *UserServiceImpl) SaveUser(user *entity.User) (*entity.User, error) {
	existingUser, _ := dal.SelectFromUsersByEmail(user.Email)
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	newUser, err := dal.InsertIntoUsers(user)

	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (userRepo *UserServiceImpl) GetUserById(id uint64) (*entity.User, error) {
	user, err := dal.SelectFromUsersById(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (userRepo *UserServiceImpl) GetUserByEmail(email string) (*entity.User, error) {
	user, err := dal.SelectFromUsersByEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil
}
