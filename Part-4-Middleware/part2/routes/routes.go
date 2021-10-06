package routes

import (
	"part2/consts"
	c "part2/controller"
	m "part2/model"
	u "part2/util"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	// user routing
	db := u.InitDB()
	mdl := m.NewUserModel(db)
	bc := c.NewController(consts.JWT_SECRET, mdl)

	e.POST("/login", bc.Login)
	e.GET("/books", bc.GetBooksController)
	e.GET("/books/:id", bc.GetBookController)

	eAuth := e.Group("")
	eAuth.Use(middleware.JWT([]byte(consts.JWT_SECRET)))
	eAuth.POST("/books", bc.CreateBookController)
	eAuth.DELETE("/books/:id", bc.DeleteBookController)
	eAuth.PUT("/books/:id", bc.UpdateBookController)

	return e
}
