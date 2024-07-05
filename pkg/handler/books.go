package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/dr4g0n369/libraryManagement/pkg/helper"
	"github.com/dr4g0n369/libraryManagement/pkg/models"
	"github.com/dr4g0n369/libraryManagement/pkg/types"
	"github.com/dr4g0n369/libraryManagement/pkg/views"
)

func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	t := views.AdminHomePage()
	if r.Method != http.MethodPost {
		t.Execute(w, types.Data{Page: "addbook"})
		return
	}

	book := types.Book{
		Name:   r.FormValue("name"),
		Author: r.FormValue("author"),
		Shelf:  r.FormValue("shelf"),
	}

	err := models.AddBook(&book)
	if err != nil {
		book.Success = -1
		log.Println(err)
	} else {
		book.Success = 1
	}

	// t.Execute(w, book)
	log.Println(book)
	t.Execute(w, types.Data{Page: "addbook", Result: book})

}

func RemoveBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusBadRequest)
		return
	}
	book := types.Book{}

	id, err := strconv.Atoi(r.FormValue("bookid"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Book Id", http.StatusBadRequest)
		return
	}

	book.BookId = uint(id)

	err = models.RemoveBook(&book)
	if err != nil {
		book.Success = -1
		log.Println(err)
	} else {
		book.Success = 1
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func AvailableBooksHandler(w http.ResponseWriter, r *http.Request) {
	t := views.HomePage()
	books, err := models.ListAllBooks()
	if err != nil {
		log.Println("Error1")
		log.Println(err)
	}

	token, _ := r.Cookie("token")

	role, err := helper.GetKey(token.Value, "role")
	if err != nil {
		log.Println(err)
	}

	log.Println(role.(string))
	if role.(string) == "admin" {
		t = views.AdminHomePage()
	}

	t.Execute(w, types.Data{Page: "availablebooks", Result: books})
}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	stringBookId := r.URL.Query().Get("bookid")

	bookid, err := strconv.Atoi(stringBookId)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	book := types.Book{
		BookId: uint(bookid),
	}

	err = models.GetBook(&book)
	if err != nil {
		book.Success = -1
	} else {
		book.Success = 1
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
