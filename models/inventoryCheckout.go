package models

import "time"

// InventoryCheckout captures info related to inventory checkout
type InventoryCheckout struct {
	CheckoutID string // unique key
	ItemID     string
	Client     string
}

// InventoryCheckoutDetails details related to checkout from inventory
type InventoryCheckoutDetails struct {
	CheckoutID          string // non-unique because checkout could be from multiple warehouses
	ItemID              string
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

func createInventoryCheckoutHistory() *InventoryCheckoutHistory {

	checkOuts := []InventoryCheckout{
		{
			CheckoutID: "1",
			ItemID:     "100",
			Client:     "Kantipur Hospital",
		},
		{
			CheckoutID: "2",
			ItemID:     "101",
			Client:     "Chaudary Group",
		},
		{
			CheckoutID: "3",
			ItemID:     "105",
			Client:     "Bhatbhateni SuperMarket",
		},
	}

	return &InventoryCheckoutHistory{
		CheckOuts: checkOuts,
	}
}
