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

	opc := 0

	for opc != 6 {
		fmt.Println("1 - Create user")
		fmt.Println("2 - List users")
		fmt.Println("3 - Get user by email")
		fmt.Println("4 - Delete user by email")
		fmt.Println("5 - Delete all users")
		fmt.Println("6 - Exit")
		fmt.Print("Option: ")
		_, err := fmt.Scanln(&opc)

		if err != nil {
			panic(err)
		}

		switch opc {
		case 1:
			var name, email string

			fmt.Print("Name: ")
			_, err := fmt.Scanln(&name)

			if err != nil {
				panic(err)
			}

			fmt.Print("Email: ")
			_, err = fmt.Scanln(&email)

			if err != nil {
				panic(err)
			}

			userCreated, _ := user.NewBuilder().WithName(name).WithEmail(email).Build()

			id, _ := userService.Create(userCreated)

			fmt.Println(id)

		case 2:
			users, err := userService.List()

			if err != nil {
				panic(err)
			}

			for _, userFromRange := range users {
				userFormatted := fmt.Sprintf("ID: %s, Name: %s, Email: %s", userFromRange.ID(), userFromRange.Name(), userFromRange.Email())

				fmt.Println(userFormatted)
			}

		case 3:
			var email string

			fmt.Print("Email: ")
			_, err := fmt.Scanln(&email)

			if err != nil {
				panic(err)
			}

			userGetByEmail, err := userService.GetByEmail(email)

			if err != nil {
				panic(err)
			}

			userFormatted := fmt.Sprintf("ID: %s, Name: %s, Email: %s", userGetByEmail.ID(), userGetByEmail.Name(), userGetByEmail.Email())

			fmt.Println(userFormatted)

		case 4:
			var email string

			fmt.Print("Email: ")
			_, err := fmt.Scanln(&email)

			if err != nil {
				panic(err)
			}

			err = userService.DeleteByEmail(email)

			if err != nil {
				panic(err)
			}

			fmt.Println("User deleted")

		case 5:
			err := userService.DeleteAll()

			if err != nil {
				panic(err)
			}

			fmt.Println("Users deleted")

		case 6:
			fmt.Println("Bye")
		default:
			fmt.Println("Invalid option")
		}
	}
}
