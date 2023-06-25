package dto

type UserElement struct {
	Id    int    `json:"id" binding:"required" example:"1"`
	Login string `json:"login" binding:"required" example:"Titouan"`
}
