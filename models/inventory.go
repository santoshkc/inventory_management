package models

import "time"

// InventoryItemDetails data related to current inventory item
type InventoryItemDetails struct {
	ItemDetailsID   string // unique key
	ItemID          string
	BatchID         string
	StockedLocation string
	StockedDate     time.Time
	ItemCount       uint
}

// Inventory current inventory status
type Inventory struct {
	Items map[string][]InventoryItemDetails
	Month string
}

// CreateInventory creates dummy inventory for functionality testing
func CreateInventory() *Inventory {
	inventoryItems := []InventoryItemDetails{
		{
			ItemDetailsID:   "1",
			ItemID:          "1",
			BatchID:         "100",
			StockedLocation: "Kathmandu",
			StockedDate:     time.Now().Local().AddDate(0, -1, 0),
			ItemCount:       1000,
		},
		{
			ItemDetailsID:   "2",
			ItemID:          "1",
			BatchID:         "100",
			StockedLocation: "Biratnagar",
			StockedDate:     time.Now().Local().AddDate(0, -2, 0),
			ItemCount:       2000,
		},
		{
			ItemDetailsID:   "3",
			ItemID:          "2",
			BatchID:         "200X",
			StockedLocation: "Chitwan",
			StockedDate:     time.Now().Local().AddDate(0, 0, -10),
			ItemCount:       3000,
		},
	}

	inventoryMap := make(map[string][]InventoryItemDetails)

	for _, itemDetail := range inventoryItems {

		val, available := inventoryMap[itemDetail.ItemID]

		if !available {
			slice := []InventoryItemDetails{itemDetail}
			inventoryMap[itemDetail.ItemID] = slice
		} else {
			slice := append(val, itemDetail)
			inventoryMap[itemDetail.ItemID] = slice
		}
	}

	return &Inventory{
		Items: inventoryMap,
		Month: "Feb",
	}
}
