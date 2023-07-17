package repository

import (
	"dip/src/core/domain/user"
	"github.com/google/uuid"
)

type UserLoader interface {
	Create(user user.User) (*uuid.UUID, error)
	List() ([]user.User, error)
}
