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

func IssueBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusBadRequest)
		return
	}
	stringBookId := r.FormValue("bookid")
	bookid, err := strconv.Atoi(stringBookId)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	book := types.Book{
		BookId: uint(bookid),
	}

	token, _ := r.Cookie("token")
	id, _ := helper.GetKey(token.Value, "id")
	/*id, err := strconv.Atoi(stringId.())
	if err != nil {
		log.Fatal(err)
	}*/

	user := types.Login{
		Id: int64(id.(float64)),
	}

	err = models.IssueBook(&book, &user)
	if err != nil {
		book.Success = -1
		log.Println(err)
		// http.Error(w, "Invalid Issue", http.StatusBadRequest)
		// return
	}

	if book.IssuedBy == user.Id {
		book.Success = 1
	}
	// t.Execute(w, book)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func ReturnBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusBadRequest)
		return
	}

	book := types.Book{}

	bookid, err := strconv.Atoi(r.FormValue("bookid"))
	if err != nil {
		log.Println(err)
		book.Success = -1
		return
	}

	book.BookId = uint(bookid)

	token, _ := r.Cookie("token")
	id, _ := helper.GetKey(token.Value, "id")
	book.IssuedBy = int64(id.(float64))

	err = models.ReturnBook(&book)
	if err != nil {
		book.Success = -1
		log.Println(err)
	} else {
		book.Success = 1
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func ListAllIssuedBooksHandler(w http.ResponseWriter, r *http.Request) {
	t := views.AdminHomePage()

	books, err := models.ListAllIssuedBooks()
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, types.Data{Page: "listallissuedbooks", Result: books})
}

func GetAllIssuedBooksByUserHandler(w http.ResponseWriter, r *http.Request) {
	t := views.HomePage()
	token, _ := r.Cookie("token")
	id, _ := helper.GetKey(token.Value, "id")

	user := types.Login{
		Id: int64(id.(float64)),
	}

	role, err := helper.GetKey(token.Value, "role")
	if err != nil {
		log.Fatal(err)
	}

	if role.(string) == "admin" {
		t = views.AdminHomePage()
	}

	books, err := models.GetIssuedBooksByUser(&user)
	if err != nil {
		log.Println(err)
	}

	t.Execute(w, types.Data{Page: "issuedbooks", Result: books})
}
