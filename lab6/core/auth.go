package core

type User struct {
	ID       int
	Username string
	Password string
	Role     string
}

type Resource struct {
	Name   string
	Access map[string]bool
}

type AuthService interface {
	Authenticate(username, password string) (User, error)
	AccessResource(user User, resourceName string) error
}
