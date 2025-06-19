package repository

import (
    "gorm.io/gorm"
    "your/module/models"
)

type GuildBattleMattingsRepository interface {
    GetByID(id string) (*models.GuildBattleMattings, error)
    Create(item *models.GuildBattleMattings) error
    Update(item *models.GuildBattleMattings) error
    Delete(id string) error
}

type guildBattleMattingsRepositoryImpl struct {
    db *gorm.DB
}

func NewGuildBattleMattingsRepository(db *gorm.DB) GuildBattleMattingsRepository {
    return &guildBattleMattingsRepositoryImpl{db: db}
}

// GORM実装
func (r *guildBattleMattingsRepositoryImpl) GetByID(id string) (*models.GuildBattleMattings, error) {
    var item models.GuildBattleMattings
    if err := r.db.First(&item, id).Error; err != nil {
        return nil, err
    }
    return &item, nil
}

func (r *guildBattleMattingsRepositoryImpl) Create(item *models.GuildBattleMattings) error {
    return r.db.Create(item).Error
}

func (r *guildBattleMattingsRepositoryImpl) BulkCreate(items ...*models.GuildBattleMattings) error {
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

func (r *guildBattleMattingsRepositoryImpl) Update(item *models.GuildBattleMattings) error {
    return r.db.Save(item).Error
}

func (r *guildBattleMattingsRepositoryImpl) BulkUpdate(items ...*models.GuildBattleMattings) error {
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

func (r *guildBattleMattingsRepositoryImpl) Delete(id string) error {
    return r.db.Delete(&models.GuildBattleMattings{}, id).Error
}
