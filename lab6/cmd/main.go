package main

import (
	"fmt"

	"github.com/ShipIM/information-security/lab6/core"
	"github.com/ShipIM/information-security/lab6/internal/service"
)

var users = []core.User{
	{ID: 1, Username: "admin", Password: "admin123", Role: "admin"},
	{ID: 2, Username: "user", Password: "user", Role: "user"},
}

var resources = []core.Resource{
	{Name: "Resource", Access: map[string]bool{"admin": true, "user": false}},
}

func main() {
	authService := service.New(users, resources)

	currentUser, err := authService.Authenticate("admin", "admin123")
	if err != nil {
		fmt.Println("Ошибка аутентификации:", err)

		return
	}

	fmt.Printf("Аутентификация успешна. Текущий пользователь: %s (Роль: %s)\n", currentUser.Username, currentUser.Role)

	fmt.Printf("Пользователь %s получил доступ к ресурсу\n", currentUser.Username)

	if err = authService.AccessResource(currentUser, "Resource"); err != nil {
		fmt.Println("Ошибка доступа:", err)
	}

	currentUser, err = authService.Authenticate("user", "user")
	if err != nil {
		fmt.Println("Ошибка аутентификации:", err)

		return
	}

	fmt.Printf("Аутентификация успешна. Текущий пользователь: %s (Роль: %s)\n", currentUser.Username, currentUser.Role)

	if err = authService.AccessResource(currentUser, "Resource"); err != nil {
		fmt.Println("Ошибка доступа:", err)
	}
}
