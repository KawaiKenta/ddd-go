package applicationservice

import (
	"errors"

	domainservice "github.com/KawaiKenta/ddd-go/domain_service"
	"github.com/KawaiKenta/ddd-go/entity"
	"github.com/KawaiKenta/ddd-go/repository"
	valueobject "github.com/KawaiKenta/ddd-go/value-object"
)

type UserApplicationService struct {
	repo repository.IUserRepository
	sv   domainservice.IUserService
}

func NewUserApplicationService(userRepository repository.IUserRepository, userService domainservice.IUserService) *UserApplicationService {
	return &UserApplicationService{
		repo: userRepository,
		sv:   userService,
	}
}

func (uas *UserApplicationService) Register(name string, email string) (*entity.User, error) {
	userName, err := valueobject.NewUserName(name)
	if err != nil {
		return nil, err
	}
	userEmail, err := valueobject.NewEmail(email)
	if err != nil {
		return nil, err
	}
	user := entity.NewUser(0, userName, userEmail)
	if uas.sv.Exists(user) {
		return nil, errors.New("user already exists")
	}

	user, err = uas.repo.Save(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
