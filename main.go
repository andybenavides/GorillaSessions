package main

import(
	"net/http"
	"github.com/gorilla/sessions"
	"html/template"
)

var tpl *template.Template

var store = sessions.NewCookieStore([]byte("secretssecrets"))

func init() {
	http.HandleFunc("/", myHandler)
	http.HandleFunc("/logout", logoutHandle)
	tpl = template.Must(template.ParseGlob("*.html"))
}

func myHandler(res http.ResponseWriter, req *http.Request){
	session, err := store.Get(req, "mySession")
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	session.Values["LoggedIn"] = true
	store.MaxAge(15)
	session.Save(req,res)

	tpl.ExecuteTemplate(res, "main.html", session)
}

func logoutHandle(res http.ResponseWriter, req *http.Request){
	session, err := store.Get(req, "mySession")
	if err != nil{
		http.Error(res, err.Error(), 500)
		return
	}

	session.Values["LoggedIn"] = false
	session.Save(req, res)
	tpl.ExecuteTemplate(res, "main.html", session)
}