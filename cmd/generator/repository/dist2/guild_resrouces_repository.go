package repository

import (
    "gorm.io/gorm"
    "your/module/models"
)

type GuildResroucesRepository interface {
    GetByID(id string) (*models.GuildResrouces, error)
    Create(item *models.GuildResrouces) error
    Update(item *models.GuildResrouces) error
    Delete(id string) error
}

type guildResroucesRepositoryImpl struct {
    db *gorm.DB
}

func NewGuildResroucesRepository(db *gorm.DB) GuildResroucesRepository {
    return &guildResroucesRepositoryImpl{db: db}
}

// GORM実装
func (r *guildResroucesRepositoryImpl) GetByID(id string) (*models.GuildResrouces, error) {
    var item models.GuildResrouces
    if err := r.db.First(&item, id).Error; err != nil {
        return nil, err
    }
    return &item, nil
}

func (r *guildResroucesRepositoryImpl) Create(item *models.GuildResrouces) error {
    return r.db.Create(item).Error
}

func (r *guildResroucesRepositoryImpl) BulkCreate(items ...*models.GuildResrouces) error {
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

func (r *guildResroucesRepositoryImpl) Update(item *models.GuildResrouces) error {
    return r.db.Save(item).Error
}

func (r *guildResroucesRepositoryImpl) BulkUpdate(items ...*models.GuildResrouces) error {
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

func (r *guildResroucesRepositoryImpl) Delete(id string) error {
    return r.db.Delete(&models.GuildResrouces{}, id).Error
}
