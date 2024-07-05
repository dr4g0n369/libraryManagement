package models

import (
	"fmt"
	"log"

	"github.com/dr4g0n369/libraryManagement/pkg/types"
)

func IssueBook(book *types.Book, user *types.Login) error {
	db, err := ConnectDatabase()
	if err != nil {
		return err
	}

	result, err := db.Exec(`INSERT INTO issued (bookid, id) VALUES (?, ?)`, book.BookId, user.Id)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	book.IssuedBy = user.Id
	return nil
}

func ReturnBook(book *types.Book) error {
	db, err := ConnectDatabase()

	if err != nil {
		log.Println(err)
		return err
	}

	result, err := db.Exec(`DELETE FROM issued WHERE bookid = ? and id = ?`, book.BookId, book.IssuedBy)
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
		return fmt.Errorf("book was not issued")
	}

	book.IssuedBy = 0
	return nil
}

func ListAllIssuedBooks() ([]types.Book, error) {
	db, err := ConnectDatabase()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	rows, err := db.Query(`SELECT books.bookid, name, author, issued.id FROM books, issued WHERE books.bookid = issued.bookid`) // check err
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.Close()
	var books []types.Book
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.BookId, &book.Name, &book.Author, &book.IssuedBy) // check err
		if err != nil {
			log.Println(err)
		} else {
			books = append(books, book)
		}
	}
	err = rows.Err() // check err
	return books, err
}

func GetIssuedBooksByUser(user *types.Login) ([]types.Book, error) {
	db, err := ConnectDatabase()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	rows, err := db.Query(`SELECT books.bookid, name, author, issued.id FROM books, issued WHERE books.bookid = issued.bookid and issued.id = ?`, user.Id) // check err
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.Close()
	var books []types.Book
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.BookId, &book.Name, &book.Author, &book.IssuedBy) // check err
		if err != nil {
			log.Println(err)
		} else {
			books = append(books, book)
		}
	}
	err = rows.Err() // check err
	return books, err
}
