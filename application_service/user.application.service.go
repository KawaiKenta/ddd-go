package applicationservice

import (
	"errors"

	domainservice "github.com/KawaiKenta/ddd-go/domain_service"
	"github.com/KawaiKenta/ddd-go/entity"
	"github.com/KawaiKenta/ddd-go/repository"
	valueobject "github.com/KawaiKenta/ddd-go/value_object"
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
	user := entity.NewUser(userName, userEmail)
	if uas.sv.Exists(user) {
		return nil, errors.New("user already exists")
	}

	user, err = uas.repo.Save(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uas *UserApplicationService) Delete(id int) error {
	// idが存在するかチェック
	user, err := uas.repo.FindById(id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	// 削除
	return uas.repo.Delete(user)
}

func (uas *UserApplicationService) Update(id int, name string, email string) (*entity.User, error) {
	// idが存在するかチェック
	user, err := uas.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// emailとnameを更新
	m, err := valueobject.NewEmail(email)
	if err != nil {
		return nil, err
	}
	n, err := valueobject.NewUserName(name)
	if err != nil {
		return nil, err
	}
	user.ChangeEmail(m)
	user.ChangeName(n)
	// emailが重複していないかチェック
	if uas.sv.Exists(user) {
		return nil, errors.New("user already exists")
	}

	// 更新
	user, err = uas.repo.Save(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
