package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dr4g0n369/libraryManagement/pkg/handler"
	"github.com/dr4g0n369/libraryManagement/pkg/middleware"
	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

func redirectToHome(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", http.StatusPermanentRedirect)
}

func main() {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	r.Use(middleware.LoggingMiddleware)

	r.HandleFunc("/register", handler.RegisterHandler)
	r.HandleFunc("/login", handler.LoginHandler)
	r.HandleFunc("/logout", handler.LogoutHandler)
	r.HandleFunc("/", redirectToHome)

	home := r.PathPrefix("/home").Subrouter()
	home.Use(middleware.AuthenticationMiddleware)

	home.HandleFunc("", handler.HomePageHandler)
	home.HandleFunc("/availablebooks", handler.AvailableBooksHandler)
	home.HandleFunc("/getbook", handler.GetBookHandler)
	home.HandleFunc("/issuebook", handler.IssueBookHandler)
	home.HandleFunc("/returnbook", handler.ReturnBookHandler)
	home.HandleFunc("/issuedbooks", handler.GetAllIssuedBooksByUserHandler)

	admin := home.PathPrefix("/admin").Subrouter()
	admin.Use(middleware.CheckAdminMiddleware)

	admin.HandleFunc("/addbook", handler.AddBookHandler)
	admin.HandleFunc("/removebook", handler.RemoveBookHandler)
	admin.HandleFunc("/allissuedbooks", handler.ListAllIssuedBooksHandler)
	admin.HandleFunc("/getuserdetails", handler.GetUserDetailsHandler)

	http.Handle("/", r)

	port := os.Getenv("WEB_PORT")
	if port == "" {
		port = "3000" // Default Port
	}

	log.Println(port)
	fmt.Printf("Started server on http://localhost:%v ...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
