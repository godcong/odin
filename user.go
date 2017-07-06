package odin

type User interface {
	ID() string
}

type user struct {
	id string
}

func NewUser() *user {
	return new(user)
}

func (u *user) ID() string {
	return ""
}
