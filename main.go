package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	Id       string `json:"id" `
	Title    string `json:"title" `
	Author   string `json:"author" `
	Quantity int    `json:"quantity" `
}

var books = []book{
	{Id: "1", Title: "Rich Dad poor Dad", Author: "Ray Dalio", Quantity: 2},
	{Id: "2", Title: "Think and Grow Rich", Author: "Dr Grow", Quantity: 3},
	{Id: "3", Title: "Richest man of Babylon", Author: "Babylony", Quantity: 6},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.Id == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func createBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.Run("localhost:8080")
}
