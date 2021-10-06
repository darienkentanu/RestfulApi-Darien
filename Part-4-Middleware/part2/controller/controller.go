package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"part2/middlewares"
	"part2/model"

	"github.com/labstack/echo"
)

type M map[string]interface{}

type LoginInfo struct {
	Title  string `json:"title" form:"title"`
	Author string `json:"author" form:"author"`
}

// var db = model.DB

type BooksModel interface {
	GetBooks() ([]model.Book, error)
	GetBook(bookId int) (model.Book, error)
	CreateBook(c echo.Context) (model.Book, error)
	DeleteBook(bookId int) error
	UpdateBook(bookId int, c echo.Context) (model.Book, error)
	Edit(id int, newB model.Book) (model.Book, error)
	GetByTitleAndAuthor(title string, author string) (model.Book, error)
}

type BookController struct {
	model      BooksModel
	JWT_SECRET string
}

func NewController(JWT_SECRET string, m BooksModel) BookController {
	return BookController{model: m, JWT_SECRET: JWT_SECRET}
}

// login
func (uc BookController) Login(c echo.Context) error {
	loginInfo := LoginInfo{}
	c.Bind(&loginInfo)
	user, err := uc.model.GetByTitleAndAuthor(loginInfo.Title, loginInfo.Author)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	token, err := middlewares.CreateToken(int(user.ID), uc.JWT_SECRET)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	user.Token = token
	user, err = uc.model.Edit(int(user.ID), user)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

// get all users
func (bc *BookController) GetBooksController(c echo.Context) error {
	books, err := bc.model.GetBooks()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, books)
}

// get user by id
func (bc *BookController) GetBookController(c echo.Context) error {
	// your solution here
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	book, err := bc.model.GetBook(bookId)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return c.JSON(http.StatusOK, book)
}

// create new user
func (bc *BookController) CreateBookController(c echo.Context) error {
	book, err := bc.model.CreateBook(c)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return c.JSON(http.StatusOK, book)
}

// delete user by id
func (bc *BookController) DeleteBookController(c echo.Context) error {
	// your solution here
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	err = bc.model.DeleteBook(bookId)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, M{"message": "deleted"})
}

func (bc *BookController) UpdateBookController(c echo.Context) error {
	// your solution here
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id")
	}
	book, err := bc.model.UpdateBook(bookId, c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, book)
}
