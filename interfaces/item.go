package interfaces

import "finanzas/misc/models"

type ItemRepo interface {
	Update(*models.Item) error
	Create(*models.Item) (int64, error)
	FindByID(int64) (*models.Item, error)
	Find(*models.Item) ([]*models.Item, error)
}
