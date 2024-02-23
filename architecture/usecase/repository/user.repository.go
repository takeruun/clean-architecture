package repository

import (
	"app/entity"
)

type UserRepository interface {
	FindAll() ([]*entity.User, error)
}
