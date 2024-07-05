package models

import (
	"fmt"
	"log"

	"github.com/dr4g0n369/libraryManagement/pkg/types"
)

func AddBook(book *types.Book) error {
	db, err := ConnectDatabase()

	if err != nil {
		return err
	}

	result, err := db.Exec(`INSERT INTO books (name, author, shelf) VALUES (?, ?, ?)`, book.Name, book.Author, book.Shelf)
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return err
	}

	book.BookId = uint(id)
	return nil
}

func RemoveBook(book *types.Book) error {
	db, err := ConnectDatabase()

	if err != nil {
		log.Println(err)
		return err
	}

	result, err := db.Exec(`DELETE FROM books WHERE bookid = ?`, book.BookId)
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	rows, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}

	if rows == 0 {
		return fmt.Errorf("no such book found")
	}

	return nil
}

func ListAllBooks() ([]types.Book, error) {
	db, err := ConnectDatabase()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	rows, err := db.Query(`SELECT books.bookid, name, author, shelf FROM books LEFT JOIN issued ON books.bookid = issued.bookid WHERE issueid is NULL`) // check err
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.Close()
	var books []types.Book
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.BookId, &book.Name, &book.Author, &book.Shelf) // check err
		if err != nil {
			log.Println(err)
		} else {
			books = append(books, book)
		}
	}
	err = rows.Err() // check err
	return books, err
}

func GetBook(book *types.Book) error {
	db, err := ConnectDatabase()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	query := "SELECT name, author, shelf FROM books WHERE bookid = ?"
	if err := db.QueryRow(query, book.BookId).Scan(&book.Name, &book.Author, &book.Shelf); err != nil {
		book.Success = -1
		log.Println(err)
		return err
	} else {
		book.Success = 1
		return nil
	}
}
