package users

import (
	"github.com/n1try/kithub2/app/config"
	"time"
)

type User struct {
	Id        string    `form:"user" binding:"required" boltholdIndex:"Id"`
	Password  string    `form:"password" binding:"required"`
	Active    bool      `form:"" boltholdIndex:"Active"`
	Gender    string    `form:"gender" binding:"required"`
	Major     string    `form:"major" binding:"required"`
	Degree    string    `form:"degree" binding:"required"`
	CreatedAt time.Time `form:""`
}

type UserQuery struct {
	ActiveEq bool
	GenderEq string
	MajorEq  string
	DegreeEq string
}

type UserSession struct {
	Token     string
	UserId    string
	CreatedAt time.Time
	LastSeen  time.Time
}

type Login struct {
	UserId   string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type UserValidator func(s *User) bool

type UserResolver func(id string) (*User, error)

type UserSessionValidator func(s *UserSession) bool

func (s *User) IsValid(validator UserValidator) bool {
	return validator(s)
}

func (s *User) IsAdmin() bool {
	for _, a := range config.Get().Auth.Admins {
		if a == s.Id {
			return true
		}
	}
	return false
}

func (s *UserSession) IsValid(validator UserSessionValidator) bool {
	return validator(s)
}
