package service

import "qubo/qubo-backend/domain/entity"

type UserService interface {
	SaveUser(user *entity.User) (*entity.User, error)
	GetUserById(id uint64) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
}
