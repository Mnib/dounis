package users

import (
	"dounis/dto"
	"dounis/models"
)

func get_all_users(db *[]models.User) []dto.UserElement {
	users := make([]dto.UserElement, len(*db))

	for i, u := range *db {
		users[i] = dto.UserElement{
			Id:    u.Id,
			Login: u.Login,
		}
	}

	return users
}

func get_user_by_id(db *[]models.User, id int) *dto.UserElement {
	for _, u := range *db {
		if u.Id == id {
			return &dto.UserElement{
				Id:    u.Id,
				Login: u.Login,
			}
		}
	}

	return nil
}
