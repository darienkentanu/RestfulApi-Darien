package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

type M map[string]interface{}

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token    string `json:"token" form:"token"`
}

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{db: db}
}

func (um *UserModel) GetByEmailAndPassword(email string, password string) (User, error) {
	u := User{}
	err := um.db.Where("email = ? AND password = ?", email, password).First(&u).Error
	return u, err
}

func (um *UserModel) GetAll() ([]User, error) {
	var users []User
	if err := um.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (um *UserModel) Add(c echo.Context) (User, error) {
	user := User{}
	if err := c.Bind(&user); err != nil {
		return User{}, err
	}
	if err := um.db.Save(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (um *UserModel) GetOne(id int) (User, error) {
	var user User
	if err := um.db.First(&user, id).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (um *UserModel) EditOne(c echo.Context, id int) (User, error) {
	var user = User{}
	if err := um.db.First(&user, id).Error; err != nil {
		return User{}, err
	}
	if user.ID == 0 {
		return User{}, fmt.Errorf("error")
	}
	if err := c.Bind(&user); err != nil {
		return User{}, err
	}
	if err := um.db.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (um *UserModel) Edit(id int, newU User) (User, error) {
	u := User{}
	err := um.db.First(&u, id).Error
	if err != nil {
		return u, err
	}
	// u.Address = newU.Address
	u.Name = newU.Name
	u.Email = newU.Email
	u.Password = newU.Password
	u.Token = newU.Token
	err = um.db.Save(&u).Error
	return u, err
}

func (um *UserModel) DeleteOne(id int) error {
	var user User
	if err := um.db.First(&user, id).Error; err != nil {
		return err
	}
	if user.ID == 0 {
		return fmt.Errorf("error")
	}
	if err := um.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
