package webserver

import (
	"fmt"
	"log"
	"net/http"
	"time"

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

type Middleware func(http.HandlerFunc) http.HandlerFunc

func loggingMiddleware() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			f(w, r)
		}
	}
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

	logging := loggingMiddleware()

	loginRouter := router.PathPrefix("/login").Subrouter()
	loginRouter.HandleFunc("/", logging(login.LoginHandler)).Methods(http.MethodGet)

	router.HandleFunc("/loginValidate", logging(login.LoginValidateHandler)).Methods(http.MethodPost)
	loginValidateRouter := router.PathPrefix("/loginValidate/").Subrouter()
	loginValidateRouter.HandleFunc("/", logging(login.LoginValidateHandler)).Methods(http.MethodPost)

	router.HandleFunc("/mainPage", logging(mainpage.MainPageHandler)).Methods(http.MethodGet)
	mainPageRouter := router.PathPrefix("/mainPage/").Subrouter()
	mainPageRouter.HandleFunc("/", logging(mainpage.MainPageHandler)).Methods(http.MethodGet)

	router.HandleFunc("/inventory", logging(inventory.InventoryHandler)).Methods(http.MethodGet)
	inventoryRouter := router.PathPrefix("/inventory").Subrouter()
	inventoryRouter.HandleFunc("/", logging(inventory.InventoryHandler)).Methods(http.MethodGet)

	router.HandleFunc("/inventoryHistory", inventory.InventoryHistoryHandler).Methods(http.MethodGet)
	inventoryHistory := router.PathPrefix("/inventoryHistory").Subrouter()
	inventoryHistory.HandleFunc("/", logging(inventory.InventoryHistoryHandler)).Methods(http.MethodGet)

	router.HandleFunc("/items", logging(items.ItemsHandler)).Methods(http.MethodGet)
	itemsRouter := router.PathPrefix("/items").Subrouter()
	itemsRouter.HandleFunc("/", logging(items.ItemsHandler)).Methods(http.MethodGet)
	itemsRouter.HandleFunc("/{uniqueId}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		uniqueID := vars["uniqueId"]
		items.ItemDetailsHandler(w, r, uniqueID)
	}).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(addr, router))
}
