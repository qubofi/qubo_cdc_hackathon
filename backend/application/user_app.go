package application

import (
	"qubo/qubo-backend/domain/entity"
	"qubo/qubo-backend/infrastructure/service_impl"
)

type userApp struct {
	userServiceImpl *service_impl.UserServiceImpl
}

var _ UserAppInterface = &userApp{}

type UserAppInterface interface {
	SaveUser(user *entity.User) (*entity.User, error)
	GetUserById(id uint64) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
}

func NewUserApplication(userServiceImpl *service_impl.UserServiceImpl) *userApp {
	return &userApp{
		userServiceImpl: userServiceImpl,
	}
}

func (u *userApp) SaveUser(user *entity.User) (*entity.User, error) {
	err := user.HashPassword()
	if err != nil {
		return nil, err
	}
	return u.userServiceImpl.SaveUser(user)
}

func (u *userApp) GetUserById(id uint64) (*entity.User, error) {
	return u.userServiceImpl.GetUserById(id)
}

func (u *userApp) GetUserByEmail(email string) (*entity.User, error) {
	return u.userServiceImpl.GetUserByEmail(email)
}
