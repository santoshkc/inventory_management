package webserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/santoshkc89/inventory_management/webserver/inventory"
	"github.com/santoshkc89/inventory_management/webserver/login"
	"github.com/santoshkc89/inventory_management/webserver/mainpage"
)

// Server web server basic information
type Server struct {
	Address string
	Port    int32
}

// Run start web server
func (server *Server) Run() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Web server active: path :%s\n", r.URL.Path)
		http.Redirect(w, r, "/login/", http.StatusFound)
	})

	http.HandleFunc("/login/", login.LoginHandler)
	http.HandleFunc("/loginValidate", login.LoginValidateHandler)

	http.HandleFunc("/mainPage", mainpage.MainPageHandler)

	http.HandleFunc("/inventory", inventory.InventoryHandler)
	http.HandleFunc("/inventoryHistory", inventory.InventoryHistoryHandler)

	addr := fmt.Sprintf("%s:%d", server.Address, server.Port)

	log.Fatal(http.ListenAndServe(addr, nil))
}
