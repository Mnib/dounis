package auth

import (
	"dounis/appdata"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	data *appdata.AppData
}

func New(data *appdata.AppData) *AuthController {
	return &AuthController{
		data,
	}
}

func (s *AuthController) Attach(r *gin.Engine) {
	auth_group := r.Group("/auth")

	auth_group.POST("/", s.login)
	auth_group.POST("/register", s.register)
}

// @Produce json
// @Summary Login endpoint
// @Tags Auth
// @Param credentials body userLogin true "Credentials"
// @Success 200 {string} Token
// @Failure 400 {string} Error message
// @Failure 404 {string} Error message
// @Router /auth [post]
func (s *AuthController) login(c *gin.Context) {
	var user_login userLogin

	if err := c.ShouldBindJSON(&user_login); err != nil {
		c.String(http.StatusBadRequest, "Incorrect payload")
		return
	}

	if token := login(user_login, s.data.Db); token != nil {
		c.String(http.StatusOK, *token)
	} else {
		c.String(http.StatusNotFound, "No such user")
	}
}

// @Produce json
// @Summary Register a user
// @Tags Auth
// @Param userInfo body userRegister true "User info"
// @Success 200 {object} dto.UserElement Newly created user
// @Failure 400 {string} Error message
// @Router /auth/register [post]
func (s *AuthController) register(c *gin.Context) {
	var user_register userRegister

	if err := c.ShouldBindJSON(&user_register); err != nil {
		c.String(http.StatusBadRequest, "Incorrect payload")
		return
	}

	if new_user := register(user_register, s.data.Db); new_user != nil {
		c.JSON(http.StatusOK, new_user)
	} else {
		c.String(http.StatusBadRequest, "Failed to add user")
	}
}
