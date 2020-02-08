package models

import "time"

// ItemCategory category of item
type ItemCategory struct {
	Name string
}

// Item describes metadata related to inventory item
type Item struct {
	UniqueID    string
	Name        string
	Description string
	Price       uint
	Category    ItemCategory
}

// ItemDetails other metadata related to inventory item
type ItemDetails struct {
	ItemDetailsID       string // unique key
	ItemID              string
	MfgDate             time.Time
	ExpiryDate          time.Time
	BatchID             string
	Price               uint
	Manufacturer        string
	ManufacturedCountry string
	Importer            string
	ImportedCountry     string
}
