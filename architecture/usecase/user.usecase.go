package usecase

import (
	"app/usecase/repository"
)

type UserUsecase struct {
	UserRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		UserRepository: userRepository,
	}
}

func (u *UserUsecase) FindAll() ([]*repository.User, error) {
	return u.UserRepository.FindAll()
}
