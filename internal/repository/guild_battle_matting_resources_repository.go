package repository

import (
    "gorm.io/gorm"
    models "github.com/dill3102/game_server_guild/internal/domain"
)

type GuildBattleMattingResourcesRepository interface {
    GetByID(id string) (*models.GuildBattleMattingResources, error)
    Create(item *models.GuildBattleMattingResources) error
    Update(item *models.GuildBattleMattingResources) error
    Delete(id string) error
}

type guildBattleMattingResourcesRepositoryImpl struct {
    db *gorm.DB
}

func NewGuildBattleMattingResourcesRepository(db *gorm.DB) GuildBattleMattingResourcesRepository {
    return &guildBattleMattingResourcesRepositoryImpl{db: db}
}

// GORM実装
func (r *guildBattleMattingResourcesRepositoryImpl) GetByID(id string) (*models.GuildBattleMattingResources, error) {
    var item models.GuildBattleMattingResources
    if err := r.db.First(&item, id).Error; err != nil {
        return nil, err
    }
    return &item, nil
}

func (r *guildBattleMattingResourcesRepositoryImpl) Create(item *models.GuildBattleMattingResources) error {
    return r.db.Create(item).Error
}

func (r *guildBattleMattingResourcesRepositoryImpl) BulkCreate(items ...*models.GuildBattleMattingResources) error {
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

func (r *guildBattleMattingResourcesRepositoryImpl) Update(item *models.GuildBattleMattingResources) error {
    return r.db.Save(item).Error
}

func (r *guildBattleMattingResourcesRepositoryImpl) BulkUpdate(items ...*models.GuildBattleMattingResources) error {
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

func (r *guildBattleMattingResourcesRepositoryImpl) Delete(id string) error {
    return r.db.Delete(&models.GuildBattleMattingResources{}, id).Error
}
