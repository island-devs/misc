package models

// Item defines the structure of services or
// articles that can be used in an invoice or bill
type Item struct {
	ID          int64   `json:"id,omitempty" gorm:"primary_key"`
	ItemTypeID  int64   `json:"item_type_id,omitempty"`
	Description string  `json:"description,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
}
