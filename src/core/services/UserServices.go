package services

import (
	"dip/src/core/domain/user"
	"dip/src/core/interfaces/primary"
	"dip/src/core/interfaces/repository"
	"github.com/google/uuid"
)

var _ primary.UserManager = (*UserServices)(nil)

type UserServices struct {
	userRepository repository.UserLoader
}

func (instance UserServices) Create(u user.User) (*uuid.UUID, error) {
	id := uuid.New()
	userCreated, _ := user.NewBuilder().WithID(id).WithName(u.Name()).WithEmail(u.Email()).Build()

	return instance.userRepository.Create(userCreated)
}

func (instance UserServices) List() ([]user.User, error) {
	return instance.userRepository.List()
}

func (instance UserServices) GetByEmail(email string) (*user.User, error) {
	return instance.userRepository.GetByEmail(email)
}

func (instance UserServices) DeleteByEmail(email string) error {
	return instance.userRepository.DeleteByEmail(email)
}

func (instance UserServices) DeleteAll() error {
	return instance.userRepository.DeleteAll()
}

func NewUserServices(userRepository repository.UserLoader) *UserServices {
	return &UserServices{userRepository: userRepository}
}
