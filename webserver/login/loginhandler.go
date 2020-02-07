package login

import (
	"errors"
	"fmt"
	"net/http"

	templateparse "github.com/santoshkc89/inventory_management/templateParse"
)

// LoginHandler used to validate login information
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Into the login page..\nUrl:%s\n", r.URL.Path)

	loginMgmt := loginManagement{
		Request:        r,
		ResponseWriter: w,
	}
	loginMgmt.handle()
}

// LoginValidateHandler this will validate whether user/pass is correct
func LoginValidateHandler(w http.ResponseWriter, r *http.Request) {

	user := r.FormValue("user")
	password := r.FormValue("pass")

	if user == userInfo.Name && password == userInfo.password {
		http.Redirect(w, r, "/mainPage", http.StatusFound)
		return
	}

	fmt.Printf("User:%s, Password: %s", user, password)

	err := errors.New("User validation error")

	http.Error(w, err.Error(), http.StatusInternalServerError)
}

// loginManagement stucture for managing the login to web server
type loginManagement struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

// handle validate login else show login page
func (loginMgmt *loginManagement) handle() {

	//loginMgmt.ResponseWriter.Header().Add("Content-type", "text/html; charset=utf-8")

	err := templateparse.RenderTemplate(loginMgmt.ResponseWriter, "login/login.html", nil)

	if err != nil {
		http.Error(loginMgmt.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}
