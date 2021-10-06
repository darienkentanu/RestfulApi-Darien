package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"part1/consts"
	"part1/middlewares"
	"part1/model"

	"github.com/labstack/echo"
)

type M map[string]interface{}

type LoginInfo struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// var DB = routes.DB

type UserModel interface {
	GetAll() ([]model.User, error)
	Add(c echo.Context) (model.User, error)
	GetOne(id int) (model.User, error)
	EditOne(c echo.Context, id int) (model.User, error)
	Edit(id int, newU model.User) (model.User, error)
	DeleteOne(id int) error
	GetByEmailAndPassword(email string, password string) (model.User, error)
}

type UserController struct {
	model      UserModel
	JWT_SECRET string
}

func (uc UserController) Login(c echo.Context) error {
	loginInfo := LoginInfo{}
	c.Bind(&loginInfo)
	user, err := uc.model.GetByEmailAndPassword(loginInfo.Email, loginInfo.Password)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "cannot login1")
	}
	token, err := middlewares.CreateToken(int(user.ID), uc.JWT_SECRET)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "cannot login2")
	}
	user.Token = token
	user, err = uc.model.Edit(int(user.ID), user)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "cannot add token")
	}
	return c.JSON(http.StatusOK, user)
}

func NewUserController(JWT_Secret string, m UserModel) UserController {
	return UserController{model: m, JWT_SECRET: consts.JWT_SECRET}
}

// get all users
func (uc UserController) GetAll(c echo.Context) error {
	users, err := uc.model.GetAll()
	if err != nil {
		return c.String(http.StatusInternalServerError, "error")
	}
	return c.JSON(http.StatusOK, M{"message": "success get all users",
		"users": users})
}

// get user by id
func (uc UserController) GetOne(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id")
	}
	user, err := uc.model.GetOne(id)
	if err != nil {
		return err
	}
	if user.ID == 0 {
		return c.String(http.StatusNotFound, "id not found")
	}
	return c.JSON(http.StatusOK, user)
}

// create new user
func (uc UserController) Add(c echo.Context) error {
	user, err := uc.model.Add(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, M{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func (uc UserController) DeleteOne(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id")
	}
	uc.model.DeleteOne(id)
	return c.JSON(http.StatusOK, M{"message": "deleted"})
}

func (uc UserController) EditOne(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, M{"message": "please input a valid id"})
	}
	user, err := uc.model.EditOne(c, id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}
