package primary

import (
	"dip/src/core/domain/user"
	"github.com/google/uuid"
)

type UserManager interface {
	Create(user user.User) (*uuid.UUID, error)
	List() ([]user.User, error)
}
