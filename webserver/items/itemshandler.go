package items

import (
	"net/http"
	"regexp"
	"time"

	"github.com/santoshkc89/inventory_management/models"
	templateparse "github.com/santoshkc89/inventory_management/templateParse"
)

// ItemRepository handles everything related to Items
type ItemRepository struct {
	items       []models.Item
	itemDetails []models.ItemDetails
}

func CreateItemRepository() *ItemRepository {

	dummyItems := func() []models.Item {
		items := []models.Item{
			{
				UniqueID:    "1",
				Name:        "Aarati Premium Rice",
				Description: "Rice that makes you want more",
				Price:       3000,
				Category: models.ItemCategory{
					Name: "Food",
				},
			},
			{
				UniqueID:    "2",
				Name:        "Wai Wai",
				Description: "Instant Noodles for instant hunger",
				Price:       20,
				Category: models.ItemCategory{
					Name: "Food",
				},
			},
			{
				UniqueID:    "3",
				Name:        "Latex Surgical Gloves",
				Description: "For hygiene and health",
				Price:       10,
				Category: models.ItemCategory{
					Name: "Health",
				},
			},
			{
				UniqueID:    "4",
				Name:        "Unique Mask",
				Description: "To prevent dust,pollution and disease",
				Price:       15,
				Category: models.ItemCategory{
					Name: "Health",
				},
			},
		}
		return items
	}

	dummyItemDetails := func() []models.ItemDetails {
		details := []models.ItemDetails{
			{
				ItemDetailsID:       "1",
				ItemID:              "1",
				MfgDate:             time.Now().AddDate(0, -3, 0),
				ExpiryDate:          time.Now().AddDate(0, 3, 0),
				BatchID:             "A1",
				Price:               3000,
				Manufacturer:        "Dugad Group",
				ManufacturedCountry: "Nepal",
				Importer:            "",
				ImportedCountry:     "",
			},
			{
				ItemDetailsID:       "2",
				ItemID:              "1",
				MfgDate:             time.Now().AddDate(0, -1, 0),
				ExpiryDate:          time.Now().AddDate(0, 5, 0),
				BatchID:             "B1",
				Price:               3500,
				Manufacturer:        "Dugad Group",
				ManufacturedCountry: "Nepal",
				Importer:            "",
				ImportedCountry:     "",
			},

			{
				ItemDetailsID:       "3",
				ItemID:              "2",
				MfgDate:             time.Now().AddDate(0, -1, 0),
				ExpiryDate:          time.Now().AddDate(0, 11, 0),
				BatchID:             "C1",
				Price:               20,
				Manufacturer:        "Chaudhary Group",
				ManufacturedCountry: "Nepal",
				Importer:            "",
				ImportedCountry:     "",
			},
			{
				ItemDetailsID:       "4",
				ItemID:              "3",
				MfgDate:             time.Now().AddDate(0, -1, 0),
				ExpiryDate:          time.Now().AddDate(0, 11, 0),
				BatchID:             "X1",
				Price:               15,
				Manufacturer:        "Ayushman Group",
				ManufacturedCountry: "India",
				Importer:            "Dirghyau Group",
				ImportedCountry:     "Nepal",
			},
			{
				ItemDetailsID:       "5",
				ItemID:              "4",
				MfgDate:             time.Now().AddDate(0, -1, 0),
				ExpiryDate:          time.Now().AddDate(0, 11, 0),
				BatchID:             "M1",
				Price:               15,
				Manufacturer:        "Nihou Equipments",
				ManufacturedCountry: "China",
				Importer:            "Healthy Nepal",
				ImportedCountry:     "Nepal",
			},
		}

		return details
	}

	items := dummyItems()
	itemDetails := dummyItemDetails()

	return &ItemRepository{
		items:       items,
		itemDetails: itemDetails,
	}
}

func (repo *ItemRepository) GetAllItemDetails() []models.ItemDetails {
	return repo.itemDetails
}

func (repo *ItemRepository) GetAllItems() []models.Item {
	return repo.items
}

func (repo *ItemRepository) GetItem(uniqueId string) (models.Item, bool) {
	for _, item := range repo.items {
		if item.UniqueID == uniqueId {
			return item, true
		}
	}
	return models.Item{}, false
}

var pathRegex *regexp.Regexp = regexp.MustCompile("^(/items)(/)?([A-Za-z0-9]+)?$")

// ItemsHandler item request handler
func ItemsHandler(w http.ResponseWriter, r *http.Request) {

	repo := CreateItemRepository()
	items := repo.GetAllItems()

	err := templateparse.RenderTemplateFromFile(w, "./webserver/items/items.html",
		"items.html", items)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// ItemDetailsHandler item request handler
func ItemDetailsHandler(w http.ResponseWriter, r *http.Request, uniqueID string) {

	repo := CreateItemRepository()
	item, _ := repo.GetItem(uniqueID)

	err := templateparse.RenderTemplateFromFile(w, "./webserver/items/itemDetails.html",
		"itemDetails.html",
		item)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
