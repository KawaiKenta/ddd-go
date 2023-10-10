package repository

import (
	"github.com/KawaiKenta/ddd-go/entity"
	valueobject "github.com/KawaiKenta/ddd-go/value-object"
)

type ID int

type IUserRepository interface {
	FindById(id int) (*entity.User, error)
	FindByEmail(email *valueobject.Email) (*entity.User, error)
	Save(user *entity.User) (*entity.User, error)
	Delete(user *entity.User) error
}

type InMemoryUserRepository struct {
	nextId ID
	users  map[ID]*entity.User
}

// IUserRepositoryを実装したInMemoryUserRepository
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		nextId: 1,
		users:  map[ID]*entity.User{},
	}
}

func (r *InMemoryUserRepository) FindById(id int) (*entity.User, error) {
	user, found := r.users[ID(id)]
	if !found {
		return nil, nil
	}
	return user, nil
}

func (r *InMemoryUserRepository) FindByEmail(email *valueobject.Email) (*entity.User, error) {
	for _, user := range r.users {
		if user.Email.Value() == email.Value() {
			return user, nil
		}
	}
	return nil, nil
}

func (r *InMemoryUserRepository) Save(user *entity.User) (*entity.User, error) {
	if user.Id == 0 {
		user.Id = int(r.nextId)
		r.nextId++
	}
	r.users[ID(user.Id)] = user
	return user, nil
}

func (r *InMemoryUserRepository) Delete(user *entity.User) error {
	delete(r.users, ID(user.Id))
	return nil
}

func (r *InMemoryUserRepository) ShowAll() {
	for _, user := range r.users {
		println("id:", user.Id, "name:", user.Name.Value(), "email:", user.Email.Value())
	}
}
