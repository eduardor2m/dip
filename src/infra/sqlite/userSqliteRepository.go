package sqlite

import (
	"context"
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

	result, err := conn.ExecContext(
		context.Background(),
		"INSERT INTO user (id, name, email) VALUES (?, ?, ?)",
		user.ID(),
		user.Name(),
		user.Email(),
	)

	if err != nil {
		return nil, err
	}

	_, err = result.LastInsertId()

	if err != nil {
		return nil, err
	}

	id := user.ID()

	return &id, nil
}

func (instance UserSqliteRepository) List() ([]user.User, error) {
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

	rows, err := conn.QueryContext(
		context.Background(),
		"SELECT id, name, email FROM user",
	)

	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()

		if err != nil {
			panic(err)
		}

	}()

	var users []user.User

	for rows.Next() {
		var id uuid.UUID
		var name string
		var email string

		err = rows.Scan(&id, &name, &email)

		if err != nil {
			return nil, err
		}

		userReturned, err := user.NewBuilder().WithID(id).WithName(name).WithEmail(email).Build()

		if err != nil {
			return nil, err
		}

		users = append(users, userReturned)
	}

	return users, nil
}

func (instance UserSqliteRepository) GetByEmail(email string) (*user.User, error) {
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

	rows, err := conn.QueryContext(
		context.Background(),
		"SELECT id, name, email FROM user WHERE email = ?",
		email,
	)

	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()

		if err != nil {
			panic(err)
		}

	}()

	var id uuid.UUID
	var name string
	var emailReturned string

	for rows.Next() {
		err = rows.Scan(&id, &name, &emailReturned)

		if err != nil {
			return nil, err
		}
	}

	userReturned, err := user.NewBuilder().WithID(id).WithName(name).WithEmail(emailReturned).Build()

	if err != nil {
		return nil, err
	}

	return &userReturned, nil
}

func (instance UserSqliteRepository) DeleteByEmail(email string) error {
	conn, err := instance.connectorManager.getConnection()

	if err != nil {
		return err
	}

	defer func() {
		err = conn.Close()

		if err != nil {
			panic(err)
		}
	}()

	_, err = conn.ExecContext(
		context.Background(),
		"DELETE FROM user WHERE email = ?",
		email,
	)

	if err != nil {
		return err
	}

	return nil
}

func (instance UserSqliteRepository) DeleteAll() error {
	conn, err := instance.connectorManager.getConnection()

	if err != nil {
		return err
	}

	defer func() {
		err = conn.Close()

		if err != nil {
			panic(err)
		}
	}()

	_, err = conn.ExecContext(
		context.Background(),
		"DELETE FROM user",
	)

	if err != nil {
		return err
	}

	return nil
}

func NewUserSqliteRepository(connectorManager connectorManager) *UserSqliteRepository {
	return &UserSqliteRepository{connectorManager: connectorManager}
}
