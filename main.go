package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"crud_api/models"
)

func main() {
	router := gin.Default()
	router.GET("/books", getBooksList)
	router.GET("/book/:id", getBookById)
	router.POST("/books", createBook)
	router.DELETE("/book/:id", deleteBookById)
	router.PUT("/book/:id", updateBookById)
	router.PATCH("/book/:id", partialUpdateBookById)
	router.Run("localhost:8080")
}

// func addBook(ctx *gin.Context) {
// 	var newBook book
// 	// Call BindJSON to bind the received JSON to
// 	// newAlbum.
// 	if err := ctx.BindJSON(&newBook); err != nil {
// 		return
// 	}

// 	// Add the new album to the slice.
// 	ctx.IndentedJSON(http.StatusCreated, newBook)
// }

func getBooksList(ctx *gin.Context) {
	products := models.GetBooks()

	if products == nil || len(products) == 0 {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.IndentedJSON(http.StatusOK, products)
	}
}

func getBookById(ctx *gin.Context) {
	id := ctx.Param("id")
	book := models.GetBook(id)

	if book == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.IndentedJSON(http.StatusOK, book)
	}
}

func createBook(ctx *gin.Context) {
	var book models.Book

	if err := ctx.BindJSON(&book); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.AddBook(book)
		ctx.IndentedJSON(http.StatusCreated, book)
	}
}

func deleteBookById(ctx *gin.Context) {
	id := ctx.Param("id")
	book := models.GetBook(id)
	if book == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		models.DeleteBook(id)
		ctx.IndentedJSON(http.StatusOK, book)
	}
}

func updateBookById(ctx *gin.Context) {
	id := ctx.Param("id")
	oldBook := models.GetBook(id)
	if oldBook == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		var updateBook models.Book
		if err := ctx.BindJSON(&updateBook); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
		} else {
			models.UpdateBook(id, updateBook)
			ctx.IndentedJSON(http.StatusOK, updateBook)
		}
	}
}

func partialUpdateBookById(ctx *gin.Context) {
	id := ctx.Param("id")
	oldBook := models.GetBook(id)
	if oldBook == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		if err := ctx.BindJSON(&oldBook); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
		} else {
			models.UpdateBook(id, *oldBook)
			ctx.IndentedJSON(http.StatusOK, oldBook)
		}
	}
}
