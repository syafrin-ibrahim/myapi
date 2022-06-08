package main

import (
	"log"
	"myapi/book"
	"myapi/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:rahasia@tcp(127.0.0.1:3306)/myapi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("koneksi putus")
	}

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)
	bookInput := book.BookInput{
		Title:       "aman",
		Description: "masih saja",
		Price:       3400,
	}

	bookService.Create(bookInput)

	// db.AutoMigrate(&book.Book{})

	//create data
	// book := book.Book{}
	// book.Title = "jujur itu mudah"
	// book.Price = 32000
	// book.Description = "dibalik jujur"
	// book.Rating = 3
	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("error created book")
	// }

	//read single data
	// var mybook book.Book
	// err = db.First(&mybook).Error
	// if err != nil {
	// 	fmt.Println("error read data book")
	// }
	// fmt.Println("judul buku = ", mybook.Title)

	//read all data
	// var mybooks []book.Book
	// err = db.Debug().Find(&mybooks).Error
	// if err != nil {
	// 	fmt.Println("error read all data book")
	// }
	// for _, b := range mybooks {
	// 	fmt.Println("judul Buku = ", b.Title)
	// }

	//update data
	// var book book.Book
	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("error read single data book")
	// }
	// book.Title = "selamat pagi"
	// book.Description = "kisah baik saat ini"
	// book.Price = 200000
	// book.Rating = 2
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("error update data book")
	// }

	//delete data
	// var book book.Book
	// err = db.Debug().Where("id = ?", 2).First(&book).Error
	// if err != nil {
	// 	fmt.Println("error read single data book")
	// }
	// err = db.Delete(&book).Error
	// if err != nil {
	// 	fmt.Println("error delete data book")
	// }
	router := gin.Default()
	//versioning
	// v1 := router.Group("/v1")
	// v1.GET("/", rootHandler)
	// v1.GET("/hello", helloHandler)
	// v1.GET("/books/:id", booksHandler)
	// v1.GET("/books", queryHandler)
	// v1.POST("/books", createBookHandler)
	router.GET("/", bookHandler.RootHandler)
	router.GET("/hello", bookHandler.HelloHandler)
	router.GET("/books/:id", bookHandler.BooksHandler)
	router.GET("/books", bookHandler.QueryHandler)
	router.POST("/books", bookHandler.CreateBookHandler)
	router.Run()
}
