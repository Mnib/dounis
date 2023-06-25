package appdata

import "dounis/models"

type AppData struct {
	Db *[]models.User
}

func New() *AppData {
	return &AppData{
		Db: &[]models.User{
			{
				Id:       1,
				Login:    "Titouan",
				Password: "LeBG",
			},
			{
				Id:       2,
				Login:    "Donatien",
				Password: "LaMenace",
			},
		},
	}
}
