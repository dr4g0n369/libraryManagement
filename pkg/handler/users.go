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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := views.LoginPage()
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	user := types.Login{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	err := models.LoginUser(&user)

	if err != nil {
		user.Success = -1
		log.Println(err)
		tmpl.Execute(w, types.Data{Page: "login", Result: user})
		return
	} else {
		user.Success = 1
		token, err := helper.CreateToken(&user)
		if err == nil {
			cookie := http.Cookie{
				Name:     "token",
				Value:    token,
				Path:     "/",
				MaxAge:   3600,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
			return
		}
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := views.RegisterPage()
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	user := types.Login{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
		Role:     "user",
	}

	if user.Username != "" && user.Password != "" {
		err := models.CreateUser(&user)
		if err != nil {
			log.Println(err)
			user.Success = -1
		} else {
			user.Success = 1
		}
	}

	tmpl.Execute(w, types.Data{Page: "register", Result: user})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
	log.Println("Cookie Set")
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := views.HomePage()
	token, err := r.Cookie("token")
	if err != nil {
		log.Println(err)
		return
	}

	role, err := helper.GetKey(token.Value, "role")
	if err != nil {
		log.Fatal(err)
	}

	if role.(string) == "admin" {
		tmpl = views.AdminHomePage()
	}

	// tmpl.Execute(w, struct{ Username string }{Username: user.(string)})
	tmpl.Execute(w, nil)
}

func GetUserDetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusBadRequest)
		return
	}
	user := types.Login{}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Book Id", http.StatusBadRequest)
		return
	}

	user.Id = int64(id)

	err = models.GetUserDetails(&user)
	if err != nil {
		user.Success = -1
		log.Println(err)
	} else {
		user.Success = 1
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
