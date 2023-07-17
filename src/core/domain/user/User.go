package user

import "github.com/google/uuid"

type User struct {
	id    uuid.UUID
	name  string
	email string
}

func (instance User) ID() uuid.UUID {
	return instance.id
}

func (instance User) Name() string {
	return instance.name
}

func (instance User) Email() string {
	return instance.email
}
