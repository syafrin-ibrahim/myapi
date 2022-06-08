package handler

import (
	"fmt"
	"myapi/book"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h bookHandler) RootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Tag": "Hello DUnia",
	})

}

func (h bookHandler) HelloHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Tag": "Hello Some One",
	})
}

func (h bookHandler) BooksHandler(ctx *gin.Context) {
	idx := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"id param ": idx,
	})
}
func (h bookHandler) QueryHandler(ctx *gin.Context) {
	title := ctx.Query("judul")
	ctx.JSON(http.StatusOK, gin.H{
		"judul = ": title,
	})
}

func (h bookHandler) CreateBookHandler(ctx *gin.Context) {
	var newBook book.BookInput

	err := ctx.ShouldBindJSON(&newBook)
	if err != nil {
		messages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("error on field %s, condition %s", e.Field(), e.ActualTag())
			messages = append(messages, errMessage)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": messages,
		})

		return
	}

	book, err := h.bookService.Create(newBook)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	response := convertToResponse(book)
	ctx.JSON(http.StatusOK, gin.H{
		"data": response,
	})

}

func convertToResponse(b book.Book) book.BookResponse {
	return book.BookResponse{
		Title:       b.Title,
		Price:       b.Price,
		Description: b.Description,
		Rating:      b.Rating,
	}
}
