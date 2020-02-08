package mainpage

import (
	"github.com/santoshkc89/inventory_management/models"
	templateparse "github.com/santoshkc89/inventory_management/templateParse"
	"github.com/santoshkc89/inventory_management/webserver/login"
	"net/http"
)

func MainPageHandler(w http.ResponseWriter, h *http.Request) {

	pageInfo := mainPageInfo{}
	pageInfo.Init()

	err := templateparse.RenderTemplate(w, "mainpage/mainpage.html", pageInfo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type mainPageInfo struct {
	User         string
	Categories   []models.ItemCategory
	ItemsInStock []models.Item
}

func (pageInfo *mainPageInfo) Init() {
	pageInfo.User = login.GetUserName()
	pageInfo.Categories = []models.ItemCategory{
		{"Medicine"},
		{"Book"},
		{"Shoe"},
		{"Computer"},
		{"Pen"},
	}
	pageInfo.ItemsInStock = []models.Item{
		models.Item{Name: "C++ How to program",
			UniqueID:    "ISBN-GUID",
			Description: "Programming related book",
			Price:       500,
			Category:    models.ItemCategory{"Book"},
		},
		models.Item{Name: "Shikhar Party Shoes",
			UniqueID:    "Shoes-GUID",
			Description: "Party shoes for men",
			Price:       1000,
			Category:    models.ItemCategory{"Shoe"},
		},
		models.Item{Name: "Toshiba Satellite Pro L50-A",
			UniqueID:    "Computer-GUID",
			Description: "Laptop for home/office use",
			Price:       100000,
			Category:    models.ItemCategory{"Computer"},
		},
		models.Item{Name: "Pantop",
			UniqueID:    "Medicine-GUID",
			Description: "Medicine related to gastrisis",
			Price:       100,
			Category:    models.ItemCategory{"Medicine"},
		},
		models.Item{Name: "Philips ball pen",
			UniqueID:    "BallPen-GUID",
			Description: "Ball pen for smooth and better writing",
			Price:       25,
			Category:    models.ItemCategory{"Pen"},
		},
	}
}

// // ItemCategory category of item
// type ItemCategory struct {
// 	Name string
// }

// // Item describes metadata related to inventory item
// type Item struct {
// 	UniqueID    string
// 	Name        string
// 	Description string
// 	Price       uint
// 	Category    ItemCategory
// }

// // ItemDetails other metadata related to inventory item
// type ItemDetails struct {
// 	ItemDetailsID       string // unique key
// 	ItemID              string
// 	MfgDate             time.Time
// 	ExpiryDate          time.Time
// 	BatchID             string
// 	Price               uint
// 	Manufacturer        string
// 	ManufacturedCountry string
// 	Importer            string
// 	ImportedCountry     string
// }

// // InventoryItemDetails data related to current inventory item
// type InventoryItemDetails struct {
// 	ItemDetailsID   string // unique key
// 	ItemID          string
// 	BatchID         string
// 	StockedLocation string
// 	StockedDate     time.Time
// 	ItemCount       uint
// }

// // Inventory current inventory status
// type Inventory struct {
// 	Items map[string]InventoryItemDetails
// }

// // InventoryCheckout captures info related to inventory checkout
// type InventoryCheckout struct {
// 	CheckoutID string // unique key
// 	ItemID     string
// 	Client     string
// }

// // InventoryCheckoutDetails details related to checkout from inventory
// type InventoryCheckoutDetails struct {
// 	CheckoutID          string // non-unique because checkout could be from multiple warehouses
// 	ItemID              string
// 	BatchID             string
// 	ItemCheckedoutCount uint
// 	CheckedoutDate      time.Time
// 	CheckedOutFrom      string
// 	CheckedOutTo        string
// }
