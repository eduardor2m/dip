package main

import (
	"dip/src/core/domain/user"
	"dip/src/core/services"
	"dip/src/infra/sqlite"
	"fmt"
)

func main() {
	userRepository := sqlite.NewUserSqliteRepository(&sqlite.DatabaseConnectionManager{})
	userService := services.NewUserServices(userRepository)

	userCreated, _ := user.NewBuilder().WithName("John Doe").WithEmail("johndoe@gmail.com").Build()

	id, _ := userService.Create(userCreated)
	users, _ := userService.List()

	fmt.Println(id)
	fmt.Println(users)
}
