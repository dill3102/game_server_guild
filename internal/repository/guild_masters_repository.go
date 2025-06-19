package repository

import (
    "gorm.io/gorm"
    models "github.com/dill3102/game_server_guild/internal/domain"
)

type GuildMastersRepository interface {
    GetByID(id string) (*models.GuildMasters, error)
    Create(item *models.GuildMasters) error
    Update(item *models.GuildMasters) error
    Delete(id string) error
}

type guildMastersRepositoryImpl struct {
    db *gorm.DB
}

func NewGuildMastersRepository(db *gorm.DB) GuildMastersRepository {
    return &guildMastersRepositoryImpl{db: db}
}

// GORM実装
func (r *guildMastersRepositoryImpl) GetByID(id string) (*models.GuildMasters, error) {
    var item models.GuildMasters
    if err := r.db.First(&item, id).Error; err != nil {
        return nil, err
    }
    return &item, nil
}

func (r *guildMastersRepositoryImpl) Create(item *models.GuildMasters) error {
    return r.db.Create(item).Error
}

func (r *guildMastersRepositoryImpl) BulkCreate(items ...*models.GuildMasters) error {
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

func (r *guildMastersRepositoryImpl) Update(item *models.GuildMasters) error {
    return r.db.Save(item).Error
}

func (r *guildMastersRepositoryImpl) BulkUpdate(items ...*models.GuildMasters) error {
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

func (r *guildMastersRepositoryImpl) Delete(id string) error {
    return r.db.Delete(&models.GuildMasters{}, id).Error
}
