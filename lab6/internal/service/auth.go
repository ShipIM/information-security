package service

import (
	"errors"

	"github.com/ShipIM/information-security/lab6/core"
)

type authService struct {
	users     []core.User
	resources []core.Resource
}

func New(users []core.User, resources []core.Resource) core.AuthService {
	return &authService{
		users:     users,
		resources: resources,
	}
}

func (as *authService) Authenticate(username, password string) (core.User, error) {
	for _, user := range as.users {
		if user.Username == username && user.Password == password {
			return user, nil
		}
	}

	return core.User{}, errors.New("неверное имя пользователя или пароль")
}

func (as *authService) AccessResource(user core.User, resourceName string) error {
	for _, resource := range as.resources {
		if resource.Name == resourceName {
			if access, ok := resource.Access[user.Role]; ok && access {
				return nil
			} else {
				return errors.New("доступ запрещен")
			}
		}
	}

	return errors.New("ресурс не найден")
}
