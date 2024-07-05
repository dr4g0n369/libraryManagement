package views

import "html/template"

func LoginPage() *template.Template {
	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	return tmpl
}

func RegisterPage() *template.Template {
	tmpl := template.Must(template.ParseFiles("templates/register.html"))
	return tmpl
}

func HomePage() *template.Template {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	return tmpl
}

func AdminHomePage() *template.Template {
	tmpl := template.Must(template.ParseFiles("templates/admin.html"))
	return tmpl
}
