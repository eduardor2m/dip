package user

import "github.com/google/uuid"

type Builder struct {
	user User
	err  error
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (instance *Builder) WithID(id uuid.UUID) *Builder {
	instance.user.id = id
	return instance
}

func (instance *Builder) WithName(name string) *Builder {
	instance.user.name = name
	return instance
}

func (instance *Builder) WithEmail(email string) *Builder {
	instance.user.email = email
	return instance
}

func (instance *Builder) Build() (User, error) {
	return instance.user, instance.err
}
