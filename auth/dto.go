package auth

type userLogin struct {
	Login    string `json:"login" binding:"required" example:"Titouan"`
	Password string `json:"password" binding:"required" example:"LeBG"`
}

type userRegister struct {
	Login    string `json:"login" binding:"required" example:"Titouan"`
	Password string `json:"password" binding:"required" example:"LeBG"`
}
