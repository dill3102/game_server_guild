package repository

import (
    "gorm.io/gorm"
    models "github.com/dill3102/game_server_guild/internal/domain"
)

type EventMastersRepository interface {
    GetByID(id string) (*models.EventMasters, error)
    Create(item *models.EventMasters) error
    Update(item *models.EventMasters) error
    Delete(id string) error
}

type eventMastersRepositoryImpl struct {
    db *gorm.DB
}

func NewEventMastersRepository(db *gorm.DB) EventMastersRepository {
    return &eventMastersRepositoryImpl{db: db}
}

// GORM実装
func (r *eventMastersRepositoryImpl) GetByID(id string) (*models.EventMasters, error) {
    var item models.EventMasters
    if err := r.db.First(&item, id).Error; err != nil {
        return nil, err
    }
    return &item, nil
}

func (r *eventMastersRepositoryImpl) Create(item *models.EventMasters) error {
    return r.db.Create(item).Error
}

func (r *eventMastersRepositoryImpl) BulkCreate(items ...*models.EventMasters) error {
    if len(items) == 0 {
        return nil
    }

    return r.db.Transaction(func(tx *gorm.DB) error {
        if err := tx.Create(items).Error; err != nil {
            return err
        }
        return nil
    })
}

func (r *eventMastersRepositoryImpl) Update(item *models.EventMasters) error {
    return r.db.Save(item).Error
}

func (r *eventMastersRepositoryImpl) BulkUpdate(items ...*models.EventMasters) error {
    if len(items) == 0 {
        return nil
    }

    return r.db.Transaction(func(tx *gorm.DB) error {
        if err := tx.Save(items).Error; err != nil {
            return err
        }
        return nil
    })
}

func (r *eventMastersRepositoryImpl) Delete(id string) error {
    return r.db.Delete(&models.EventMasters{}, id).Error
}
