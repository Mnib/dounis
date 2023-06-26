package auth

import (
	"dounis/dto"
	"dounis/models"
	"fmt"
)

func login(credentials userLogin, db *[]models.User) *string {
	fmt.Printf("User wants to login with login %s and password %s\n", credentials.Login, credentials.Password)

	for _, u := range *db {
		if u.Login == credentials.Login && u.Password == credentials.Password {
			if token, err := getTokenForUser(Claims{Id: u.Id}); err != nil {
				return nil
			} else {
				return &token
			}
		}
	}

	return nil
}

func register(userInfo userRegister, db *[]models.User) *dto.UserElement {
	fmt.Printf("User wants to register with login %s and password %s\n", userInfo.Login, userInfo.Password)

	new_user := models.User{
		Id:       len(*db) + 1,
		Login:    userInfo.Login,
		Password: userInfo.Password,
	}

	*db = append(*db, new_user)

	return &dto.UserElement{
		Id:    new_user.Id,
		Login: new_user.Login,
	}
}
