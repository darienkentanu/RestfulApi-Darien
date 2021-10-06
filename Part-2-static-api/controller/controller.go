package controller

import (
	"net/http"
	"strconv"

	"github.com/darienkentanu/RESTful-API-with-Go/model"
	"github.com/labstack/echo"
)

type M map[string]interface{}

var users []model.User

// ----------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, M{
		"messages": "success get all users",
		"users":    users,
	})
}

func contains(s []model.User, idx int) bool {
	for i := range s {
		if i == idx {
			return true
		}
	}
	return false
}

func GetUserController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, M{"messages": "please input a valid routes"})
	}
	if contains(users, id-1) {
		return c.JSON(http.StatusOK, M{
			"messages": "success get user",
			"user":     users[id-1],
		})
	}
	return c.JSON(http.StatusBadRequest, M{"message": "id doesn't exist"})
}

func DeleteUserController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, M{"message": "please input a valid routes"})
	}
	if contains(users, id-1) {
		users = append(users[:id-1], users[id:]...)
		return c.JSON(http.StatusOK, M{"message": "success to delete user"})
	}
	return c.JSON(http.StatusBadRequest, M{"message": "id doesn't exist"})
}

func UpdateUserControllers(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, M{"message": "please input a valid id"})
	}
	if contains(users, id-1) {
		user := &users[id-1]
		user.Name = c.FormValue("name")
		user.Email = c.FormValue("email")
		user.Password = c.FormValue("password")
		return c.JSON(http.StatusOK, M{
			"messages":     "success",
			"Updated user": user,
		})
	}
	return c.JSON(http.StatusBadRequest, M{"message": "id doesn't exist"})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := model.User{}
	user.Name = c.FormValue("name")
	if len(user.Name) == 0 {
		return c.JSON(http.StatusBadRequest, M{"messages": "please input a valid username"})
	}
	user.Email = c.FormValue("email")
	if len(user.Email) == 0 {
		return c.JSON(http.StatusBadRequest, M{"messages": "please input a valid email"})
	}
	user.Password = c.FormValue("password")
	if len(user.Password) == 0 {
		return c.JSON(http.StatusBadRequest, M{"messages": "please input a valid password"})
	}
	// binding data
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, M{
		"messages": "success create user",
		"user":     user,
	})
}
