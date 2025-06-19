package repository

import (
    "gorm.io/gorm"
    models "github.com/dill3102/game_server_guild/internal/domain"
)

type GuildMembersRepository interface {
    GetByID(id string) (*models.GuildMembers, error)
    Create(item *models.GuildMembers) error
    Update(item *models.GuildMembers) error
    Delete(id string) error
}

type guildMembersRepositoryImpl struct {
    db *gorm.DB
}

func NewGuildMembersRepository(db *gorm.DB) GuildMembersRepository {
    return &guildMembersRepositoryImpl{db: db}
}

// GORM実装
func (r *guildMembersRepositoryImpl) GetByID(id string) (*models.GuildMembers, error) {
    var item models.GuildMembers
    if err := r.db.First(&item, id).Error; err != nil {
        return nil, err
    }
    return &item, nil
}

func (r *guildMembersRepositoryImpl) Create(item *models.GuildMembers) error {
    return r.db.Create(item).Error
}

func (r *guildMembersRepositoryImpl) BulkCreate(items ...*models.GuildMembers) error {
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

func (r *guildMembersRepositoryImpl) Update(item *models.GuildMembers) error {
    return r.db.Save(item).Error
}

func (r *guildMembersRepositoryImpl) BulkUpdate(items ...*models.GuildMembers) error {
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

func (r *guildMembersRepositoryImpl) Delete(id string) error {
    return r.db.Delete(&models.GuildMembers{}, id).Error
}
