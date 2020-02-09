package inventory

import (
	"net/http"

	"github.com/santoshkc89/inventory_management/models"

	templateparse "github.com/santoshkc89/inventory_management/templateParse"
)

// import "github.com/santoshkc89/inventory_management/models"

func InventoryHandler(w http.ResponseWriter, r *http.Request) {

	inventory := models.CreateInventory()

	err := templateparse.RenderTemplate(w, "inventory/inventory.html", *inventory)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func InventoryHistoryHandler(w http.ResponseWriter, r *http.Request) {

	inventoryHistory := models.CreateInventoryCheckoutHistory()

	err := templateparse.RenderTemplateFromFile(w, "./webserver/inventory/inventoryHistory.html",
		"inventoryHistory.html", inventoryHistory)

	//err := templateparse.RenderTemplate(w, "inventory/inventoryHistory.html", *inventoryHistory)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
