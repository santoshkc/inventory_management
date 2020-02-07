package mainpage

import (
	"net/http"

	templateparse "github.com/santoshkc89/inventory_management/templateParse"
	"github.com/santoshkc89/inventory_management/webserver/login"
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
	Categories   []ItemCategory
	ItemsInStock []Item
}

func (pageInfo *mainPageInfo) Init() {
	pageInfo.User = login.GetUserName()
	pageInfo.Categories = []ItemCategory{
		{"Medicine"},
		{"Book"},
		{"Shoe"},
		{"Computer"},
		{"Pen"},
	}
	pageInfo.ItemsInStock = []Item{
		Item{Name: "C++ How to program",
			UniqueID:    "ISBN-GUID",
			Description: "Programming related book",
			Price:       500,
			Category:    ItemCategory{"Book"},
		},
		Item{Name: "Shikhar Party Shoes",
			UniqueID:    "Shoes-GUID",
			Description: "Party shoes for men",
			Price:       1000,
			Category:    ItemCategory{"Shoe"},
		},
		Item{Name: "Toshiba Satellite Pro L50-A",
			UniqueID:    "Computer-GUID",
			Description: "Laptop for home/office use",
			Price:       100000,
			Category:    ItemCategory{"Computer"},
		},
		Item{Name: "Pantop",
			UniqueID:    "Medicine-GUID",
			Description: "Medicine related to gastrisis",
			Price:       100,
			Category:    ItemCategory{"Medicine"},
		},
		Item{Name: "Philips ball pen",
			UniqueID:    "BallPen-GUID",
			Description: "Ball pen for smooth and better writing",
			Price:       25,
			Category:    ItemCategory{"Pen"},
		},
	}
}

// ItemCategory catogory of item
type ItemCategory struct {
	Name string
}

// Item describes metadata related to inventory item
type Item struct {
	Name        string
	UniqueID    string
	Description string
	Price       uint
	Category    ItemCategory
}
