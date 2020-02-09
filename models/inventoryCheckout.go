package models

import (
	"time"
)

// InventoryCheckout captures info related to inventory checkout
type InventoryCheckout struct {
	CheckoutID      string // unique key
	Client          string
	CheckoutDetails []InventoryCheckoutDetails
}

// InventoryCheckoutDetails details related to checkout from inventory
type InventoryCheckoutDetails struct {
	CheckoutID          string // non-unique because checkout could be from multiple warehouses
	ItemID              string
	ItemName            string
	BatchID             string
	ItemCheckedoutCount uint
	CheckedoutDate      time.Time
	CheckedOutFrom      string
	CheckedOutTo        string
}

// InventoryCheckoutHistory inventory history
type InventoryCheckoutHistory struct {
	CheckOuts []InventoryCheckout
}

func CreateInventoryCheckoutHistory() *InventoryCheckoutHistory {

	checkOuts := []InventoryCheckout{
		{
			CheckoutID: "1",
			Client:     "Bhatbhateni Supermarket",
			CheckoutDetails: []InventoryCheckoutDetails{
				{
					CheckoutID:          "1",
					ItemID:              "100",
					ItemName:            "Wai Wai",
					BatchID:             "A1",
					ItemCheckedoutCount: 1000,
					CheckedoutDate:      time.Now().AddDate(0, -1, 0),
					CheckedOutFrom:      "Kathmandu",
					CheckedOutTo:        "Rasuwa",
				},
				{
					CheckoutID:          "1",
					ItemID:              "100",
					ItemName:            "Gillete Razor",
					BatchID:             "A1",
					ItemCheckedoutCount: 2000,
					CheckedoutDate:      time.Now().AddDate(0, -1, 0),
					CheckedOutFrom:      "Biratnagar",
					CheckedOutTo:        "Khotang",
				},
			},
		},
		{
			CheckoutID: "2",
			Client:     "National Trading",
			CheckoutDetails: []InventoryCheckoutDetails{
				{
					CheckoutID:          "2",
					ItemID:              "100",
					ItemName:            "Aarati Premium Rice",
					BatchID:             "A1",
					ItemCheckedoutCount: 1000,
					CheckedoutDate:      time.Now().AddDate(0, 0, -15),
					CheckedOutFrom:      "Pokhara",
					CheckedOutTo:        "Lamjung",
				},
				{
					CheckoutID:          "2",
					ItemID:              "101",
					ItemName:            "Jumla Apple",
					BatchID:             "B1",
					ItemCheckedoutCount: 2000,
					CheckedoutDate:      time.Now().AddDate(0, 0, -15),
					CheckedOutFrom:      "Nepalgunj",
					CheckedOutTo:        "Mugu",
				},
			},
		},
		{
			CheckoutID: "3",
			Client:     "Kathmandu Medical College",
			CheckoutDetails: []InventoryCheckoutDetails{
				{
					CheckoutID:          "3",
					ItemID:              "103",
					ItemName:            "Surgical Gloves",
					BatchID:             "C1",
					ItemCheckedoutCount: 1000,
					CheckedoutDate:      time.Now().AddDate(0, 0, -7),
					CheckedOutFrom:      "Pokhara",
					CheckedOutTo:        "Kathmandu",
				},
				{
					CheckoutID:          "3",
					ItemID:              "104",
					ItemName:            "Pantop",
					BatchID:             "D1",
					ItemCheckedoutCount: 10000,
					CheckedoutDate:      time.Now().AddDate(0, 0, -7),
					CheckedOutFrom:      "Biratnagar",
					CheckedOutTo:        "Kathmandu",
				},
			},
		},
	}

	return &InventoryCheckoutHistory{
		CheckOuts: checkOuts,
	}
}
