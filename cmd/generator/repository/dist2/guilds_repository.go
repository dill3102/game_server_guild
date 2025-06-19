package repository

import (
    "gorm.io/gorm"
    "your/module/models"
)

type GuildsRepository interface {
    GetByID(id string) (*models.Guilds, error)
    Create(item *models.Guilds) error
    Update(item *models.Guilds) error
    Delete(id string) error
}

type guildsRepositoryImpl struct {
    db *gorm.DB
}

func NewGuildsRepository(db *gorm.DB) GuildsRepository {
    return &guildsRepositoryImpl{db: db}
}

// GORM実装
func (r *guildsRepositoryImpl) GetByID(id string) (*models.Guilds, error) {
    var item models.Guilds
    if err := r.db.First(&item, id).Error; err != nil {
        return nil, err
    }
    return &item, nil
}

func (r *guildsRepositoryImpl) Create(item *models.Guilds) error {
    return r.db.Create(item).Error
}

func (r *guildsRepositoryImpl) BulkCreate(items ...*models.Guilds) error {
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

func (r *guildsRepositoryImpl) Update(item *models.Guilds) error {
    return r.db.Save(item).Error
}

func (r *guildsRepositoryImpl) BulkUpdate(items ...*models.Guilds) error {
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

func (r *guildsRepositoryImpl) Delete(id string) error {
    return r.db.Delete(&models.Guilds{}, id).Error
}
