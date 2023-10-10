package domainservice

import (
	"github.com/KawaiKenta/ddd-go/entity"
	"github.com/KawaiKenta/ddd-go/repository"
)

type IUserService interface {
	Exists(user *entity.User) bool
}

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) *UserService {
	return &UserService{
		repo: userRepository,
	}
}

func (us *UserService) Exists(user *entity.User) bool {
	user, err := us.repo.FindByEmail(user.Email)
	if user != nil && err == nil {
		return true
	}
	return false
}
