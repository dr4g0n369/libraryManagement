package middleware

import (
	"log"
	"net/http"

	"github.com/dr4g0n369/libraryManagement/pkg/helper"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		if err != nil {
			log.Println(err)
			log.Println("Redirecting to /login")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		err = helper.VerifyToken(token.Value)
		if err != nil {
			log.Println(err)
			log.Println("Redirecting to /login")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func CheckAdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		if err != nil {
			log.Fatal(err)
		}

		role, err := helper.GetKey(token.Value, "role")
		if err != nil {
			log.Fatal(err)
		}

		if role != "admin" {
			http.Error(w, "Forbidden", http.StatusForbidden)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
