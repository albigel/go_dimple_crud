package models

import (
	"database/sql"

	_ "github.com/lib/pq"

	"fmt"
)

const (
	host     = "localhost"
	port     = 55000
	user     = "postgres"
	password = "postgrespw"
	dbname   = "postgres"
)

type Book struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	Pageamount int    `json:"pageamount"`
}

func connectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	// if there is an error opening the connection, handle it
	if err != nil {
		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}
	return db
}

func GetBooks() []Book {
	db := connectDB()
	defer db.Close()
	books := []Book{}
	results, err := db.Query("SELECT * FROM books")
	for results.Next() {
		var book Book
		// for each row, scan into the Product struct
		err = results.Scan(&book.ID, &book.Title, &book.Author, &book.Pageamount)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// append the product into products array
		books = append(books, book)
	}

	return books
}

func GetBook(id string) *Book {
	db := connectDB()
	defer db.Close()

	results, err := db.Query(`SELECT * FROM books WHERE id=$1`, id)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	book := &Book{}
	if results.Next() {
		err = results.Scan(&book.ID, &book.Title, &book.Author, &book.Pageamount)
		if err != nil {
			return nil
		}
	} else {

		return nil
	}

	return book
}

func AddBook(book Book) {

	db := connectDB()
	defer db.Close()

	insert, err := db.Query(
		`INSERT INTO books (author,title,pageamount) VALUES ($1,$2,$3)`,
		book.Author, book.Title, book.Pageamount)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}

func DeleteBook(id string) {
	db := connectDB()
	defer db.Close()

	_, err := db.Query(`DELETE FROM books WHERE id=$1`, id)

	if err != nil {
		fmt.Println("Err", err.Error())
		panic(err.Error())
	}
}

func UpdateBook(id string, book Book) {
	db := connectDB()
	defer db.Close()
	_, err := db.Query(`UPDATE books SET author=$2,title=$3,pageamount=$4 WHERE id=$1`, id, book.Author, book.Title, book.Pageamount)
	if err != nil {
		fmt.Println("Err", err.Error())
		panic(err.Error())
	}
}
