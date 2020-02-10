package webserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/santoshkc89/inventory_management/webserver/inventory"
	"github.com/santoshkc89/inventory_management/webserver/items"
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

	// handle static resource
	fsHandler := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("static/", fsHandler))

	addr := fmt.Sprintf("%s:%d", server.Address, server.Port)

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusFound)
	}).Methods(http.MethodGet)

	loginRouter := router.PathPrefix("/login").Subrouter()
	loginRouter.HandleFunc("/", login.LoginHandler).Methods(http.MethodGet)

	router.HandleFunc("/loginValidate", login.LoginValidateHandler).Methods(http.MethodPost)
	loginValidateRouter := router.PathPrefix("/loginValidate/").Subrouter()
	loginValidateRouter.HandleFunc("/", login.LoginValidateHandler).Methods(http.MethodPost)

	router.HandleFunc("/mainPage", mainpage.MainPageHandler).Methods(http.MethodGet)
	mainPageRouter := router.PathPrefix("/mainPage/").Subrouter()
	mainPageRouter.HandleFunc("/", mainpage.MainPageHandler).Methods(http.MethodGet)

	router.HandleFunc("/inventory", inventory.InventoryHandler).Methods(http.MethodGet)
	inventoryRouter := router.PathPrefix("/inventory").Subrouter()
	inventoryRouter.HandleFunc("/", inventory.InventoryHandler).Methods(http.MethodGet)

	router.HandleFunc("/inventoryHistory", inventory.InventoryHistoryHandler).Methods(http.MethodGet)
	inventoryHistory := router.PathPrefix("/inventoryHistory").Subrouter()
	inventoryHistory.HandleFunc("/", inventory.InventoryHistoryHandler).Methods(http.MethodGet)

	router.HandleFunc("/items", items.ItemsHandler).Methods(http.MethodGet)
	itemsRouter := router.PathPrefix("/items").Subrouter()
	itemsRouter.HandleFunc("/", items.ItemsHandler).Methods(http.MethodGet)
	itemsRouter.HandleFunc("/{uniqueId}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		uniqueID := vars["uniqueId"]
		items.ItemDetailsHandler(w, r, uniqueID)
	}).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(addr, router))
}
