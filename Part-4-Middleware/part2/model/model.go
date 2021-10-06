package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

type Book struct {
	gorm.Model
	Title  string `json:"title" form:"title"`
	Author string `json:"author" form:"author"`
	Token  string `json:"Token" form:"token"`
}

type BooksModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *BooksModel {
	return &BooksModel{db: db}
}

func (bm *BooksModel) GetByTitleAndAuthor(title string, author string) (Book, error) {
	b := Book{}
	err := bm.db.Where("title = ? AND author = ?", title, author).First(&b).Error
	return b, err
}

func (bm *BooksModel) GetBooks() ([]Book, error) {
	var books []Book
	if err := bm.db.Find(&books).Error; err != nil {
		fmt.Println(err)
		return []Book{}, err
	}
	return books, nil
}

func (bm *BooksModel) GetBook(bookId int) (Book, error) {
	var book Book
	if err := bm.db.First(&book, bookId).Error; err != nil {
		fmt.Println(err)
		return Book{}, err
	}
	if book.ID == 0 {
		return Book{}, fmt.Errorf("Book not found")
	}
	return book, nil
}

func (bm *BooksModel) CreateBook(c echo.Context) (Book, error) {
	var book = Book{}
	if err := c.Bind(&book); err != nil {
		return Book{}, err
	}
	if err := bm.db.Save(&book).Error; err != nil {
		return Book{}, err
	}
	return book, nil
}

func (bm *BooksModel) DeleteBook(bookId int) error {
	var book Book
	if err := bm.db.First(&book, bookId).Error; err != nil {
		return err
	}
	if book.ID == 0 {
		return fmt.Errorf("book not found")
	}
	if err := bm.db.Delete(&book).Error; err != nil {
		return err
	}
	return nil
}

func (bm *BooksModel) UpdateBook(bookId int, c echo.Context) (Book, error) {
	var book Book
	if err := bm.db.First(&book, bookId).Error; err != nil {
		return Book{}, err
	}
	if book.ID == 0 {
		return Book{}, fmt.Errorf("Book not Found")
	}
	if err := c.Bind(&book); err != nil {
		return Book{}, err
	}
	if err := bm.db.Save(&book).Error; err != nil {
		return Book{}, err
	}
	return book, nil
}

func (bm *BooksModel) Edit(id int, newB Book) (Book, error) {
	b := Book{}
	err := bm.db.First(&b, id).Error
	if err != nil {
		return b, err
	}
	// u.Address = newU.Address
	b.Author = newB.Author
	b.Title = newB.Title
	b.Token = newB.Token
	err = bm.db.Save(&b).Error
	return b, err
}
