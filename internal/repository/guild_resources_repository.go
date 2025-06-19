package repository

import (
    "gorm.io/gorm"
    models "github.com/dill3102/game_server_guild/internal/domain"
)

type GuildResourcesRepository interface {
    GetByID(id string) (*models.GuildResources, error)
    Create(item *models.GuildResources) error
    Update(item *models.GuildResources) error
    Delete(id string) error
}

type guildResourcesRepositoryImpl struct {
    db *gorm.DB
}

func NewGuildResourcesRepository(db *gorm.DB) GuildResourcesRepository {
    return &guildResourcesRepositoryImpl{db: db}
}

// GORM実装
func (r *guildResourcesRepositoryImpl) GetByID(id string) (*models.GuildResources, error) {
    var item models.GuildResources
    if err := r.db.First(&item, id).Error; err != nil {
        return nil, err
    }
    return &item, nil
}

func (r *guildResourcesRepositoryImpl) Create(item *models.GuildResources) error {
    return r.db.Create(item).Error
}

func (r *guildResourcesRepositoryImpl) BulkCreate(items ...*models.GuildResources) error {
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

func (r *guildResourcesRepositoryImpl) Update(item *models.GuildResources) error {
    return r.db.Save(item).Error
}

func (r *guildResourcesRepositoryImpl) BulkUpdate(items ...*models.GuildResources) error {
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

func (r *guildResourcesRepositoryImpl) Delete(id string) error {
    return r.db.Delete(&models.GuildResources{}, id).Error
}
