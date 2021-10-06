package routes

import (
	"part1/consts"
	c "part1/controller"
	m "part1/model"
	u "part1/util"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	// user routing
	DB := u.InitDB()
	userModel := m.NewUserModel(DB)
	uc := c.NewUserController(consts.JWT_SECRET, userModel)

	e.POST("/login", uc.Login)
	e.GET("/users", uc.GetAll)

	eAuth := e.Group("")
	eAuth.Use(middleware.JWT([]byte(consts.JWT_SECRET)))

	eAuth.GET("/users/:id", uc.GetOne)
	eAuth.POST("/users", uc.Add)
	eAuth.DELETE("/users/:id", uc.DeleteOne)
	eAuth.PUT("/users/:id", uc.EditOne)

	return e
}
