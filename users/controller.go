package users

import (
	"dounis/appdata"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	data *appdata.AppData
}

func New(data *appdata.AppData) *UserController {
	return &UserController{
		data,
	}
}

func (c *UserController) Attach(r *gin.Engine) {
	auth_group := r.Group("/users")

	auth_group.GET("/", c.get_all_users)
	auth_group.GET("/:id", c.get_user_by_id)
}

// @Produce json
// @Summary Get the list of registered users
// @Tags Users
// @Success 200 {array} dto.UserElement List of users
// @Failure 401 {string} Unauthorized
// @Router /users [get]
func (s *UserController) get_all_users(c *gin.Context) {
	c.JSON(http.StatusOK, get_all_users(s.data.Db))
}

// @Produce json
// @Summary Get the list of registered users
// @Tags Users
// @Param id path int true "ID of the user"
// @Success 200 {object} dto.UserElement Requested user
// @Failure 401 {string} Unauthorized
// @Failure 404 {string} NotFound
// @Router /users/{id} [get]
func (s *UserController) get_user_by_id(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	if user := get_user_by_id(s.data.Db, id); user != nil {
		c.JSON(http.StatusOK, user)
	} else {
		c.Status(http.StatusNotFound)
	}
}
