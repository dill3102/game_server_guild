package repository

import (
    "gorm.io/gorm"
    "your/module/models"
)

type GuildBattleMattingCharactersRepository interface {
    GetByID(id string) (*models.GuildBattleMattingCharacters, error)
    Create(item *models.GuildBattleMattingCharacters) error
    Update(item *models.GuildBattleMattingCharacters) error
    Delete(id string) error
}

type guildBattleMattingCharactersRepositoryImpl struct {
    db *gorm.DB
}

func NewGuildBattleMattingCharactersRepository(db *gorm.DB) GuildBattleMattingCharactersRepository {
    return &guildBattleMattingCharactersRepositoryImpl{db: db}
}

// GORM実装
func (r *guildBattleMattingCharactersRepositoryImpl) GetByID(id string) (*models.GuildBattleMattingCharacters, error) {
    var item models.GuildBattleMattingCharacters
    if err := r.db.First(&item, id).Error; err != nil {
        return nil, err
    }
    return &item, nil
}

func (r *guildBattleMattingCharactersRepositoryImpl) Create(item *models.GuildBattleMattingCharacters) error {
    return r.db.Create(item).Error
}

func (r *guildBattleMattingCharactersRepositoryImpl) BulkCreate(items ...*models.GuildBattleMattingCharacters) error {
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

func (r *guildBattleMattingCharactersRepositoryImpl) Update(item *models.GuildBattleMattingCharacters) error {
    return r.db.Save(item).Error
}

func (r *guildBattleMattingCharactersRepositoryImpl) BulkUpdate(items ...*models.GuildBattleMattingCharacters) error {
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

func (r *guildBattleMattingCharactersRepositoryImpl) Delete(id string) error {
    return r.db.Delete(&models.GuildBattleMattingCharacters{}, id).Error
}
