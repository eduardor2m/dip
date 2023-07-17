package sqlite

import (
	"dip/src/core/domain/user"
	"dip/src/core/interfaces/repository"
	"github.com/google/uuid"
)

var _ repository.UserLoader = &UserSqliteRepository{}

type UserSqliteRepository struct {
	connectorManager
}

func (instance UserSqliteRepository) Create(user user.User) (*uuid.UUID, error) {
	conn, err := instance.connectorManager.getConnection()

	if err != nil {
		return nil, err
	}

	defer func() {
		err = conn.Close()

		if err != nil {
			panic(err)
		}
	}()

	if err != nil {
		return nil, err
	}

	id := user.ID()

	return &id, nil
}

func (instance UserSqliteRepository) List() ([]user.User, error) {
	return []user.User{}, nil
}

func NewUserSqliteRepository(connectorManager connectorManager) *UserSqliteRepository {
	return &UserSqliteRepository{connectorManager: connectorManager}
}
