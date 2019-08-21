package repos

import (
	"errors"
	"finanzas/misc/models"

	"github.com/jinzhu/gorm"
)

// ItemRepo refers to a struct
// that contains all methods to
// manipulate an Item and its interactions with the DB
type ItemRepo struct {
	db *gorm.DB
}

// NewItemRepo returns a ready to use reference to an
// ItemRepo struct.
func NewItemRepo(db *gorm.DB) *ItemRepo {
	return &ItemRepo{
		db: db,
	}
}

// FindByID receives an ID as an int64 and queries the Item table
// for the first item to match the given ID.
func (i *ItemRepo) FindByID(id int64) (*models.Item, error) {

	item := &models.Item{}
	err := i.db.First(item, id).Error
	if err != nil {
		return nil, err
	}

	return item, nil
}

// Find receives a full item struct as a query (since the query can be by any field).
// It will return all records matching with the supplied parameters.
func (i *ItemRepo) Find(query *models.Item) ([]*models.Item, error) {

	items := make([]*models.Item, 0)

	err := i.db.Find(&items, query).Error
	if err != nil {
		return nil, err
	}

	return items, nil
}

// Create receives a full item struct and a new Item record will
// be created in the database with the supplied info
func (i *ItemRepo) Create(item *models.Item) (int64, error) {

	if !i.db.NewRecord(item) {
		return 0, errors.New("The item to be created must be new")
	}

	err := i.db.Create(item).Error
	if err != nil {
		return 0, err
	}

	return item.ID, nil
}

// Update receives a full item struct and an existing Item record
// will be updated. Note that the item supplied must exist in the database first.
func (i *ItemRepo) Update(item *models.Item) error {

	err := i.db.Model(item).Updates(item).Error
	if err != nil {
		return err
	}

	return nil
}
