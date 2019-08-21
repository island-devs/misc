package models

// ItemProcess defines both which accounts are affected by each item
// and also what document types can use the items
type ItemProcess struct {
	ID             int64 `json:"id,omitempty" gorm:"primary_key"`
	ItemID         int64 `json:"item_id,omitempty"`
	AccountID      int64 `json:"account_id,omitempty"`
	DocumentTypeID int64 `json:"document_type_id,omitempty"`
}
