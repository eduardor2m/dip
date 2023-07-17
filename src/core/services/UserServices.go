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

func NewUserServices(userRepository repository.UserLoader) *UserServices {
	return &UserServices{userRepository: userRepository}
}
