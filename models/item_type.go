package models

// ItemType classifies items used for invoicing and billing
type ItemType struct {
	ID          int64  `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
}
