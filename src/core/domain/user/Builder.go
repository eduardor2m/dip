package user

import (
	"errors"
	"github.com/google/uuid"
	"regexp"
)

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
	if name == "" {
		instance.err = errors.New("name is required")
		return instance
	}

	instance.user.name = name
	return instance
}

func (instance *Builder) WithEmail(email string) *Builder {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	if email == "" {
		instance.err = errors.New("email is required")
		return instance
	} else if !emailRegex.MatchString(email) {
		instance.err = errors.New("email is invalid")
		return instance
	}

	instance.user.email = email
	return instance
}

func (instance *Builder) Build() (User, error) {
	return instance.user, instance.err
}
