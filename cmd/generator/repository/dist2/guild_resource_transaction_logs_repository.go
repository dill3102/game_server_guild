package repository

import (
    "gorm.io/gorm"
    "your/module/models"
)

type GuildResourceTransactionLogsRepository interface {
    GetByID(id string) (*models.GuildResourceTransactionLogs, error)
    Create(item *models.GuildResourceTransactionLogs) error
    Update(item *models.GuildResourceTransactionLogs) error
    Delete(id string) error
}

type guildResourceTransactionLogsRepositoryImpl struct {
    db *gorm.DB
}

func NewGuildResourceTransactionLogsRepository(db *gorm.DB) GuildResourceTransactionLogsRepository {
    return &guildResourceTransactionLogsRepositoryImpl{db: db}
}

// GORM実装
func (r *guildResourceTransactionLogsRepositoryImpl) GetByID(id string) (*models.GuildResourceTransactionLogs, error) {
    var item models.GuildResourceTransactionLogs
    if err := r.db.First(&item, id).Error; err != nil {
        return nil, err
    }
    return &item, nil
}

func (r *guildResourceTransactionLogsRepositoryImpl) Create(item *models.GuildResourceTransactionLogs) error {
    return r.db.Create(item).Error
}

func (r *guildResourceTransactionLogsRepositoryImpl) BulkCreate(items ...*models.GuildResourceTransactionLogs) error {
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

func (r *guildResourceTransactionLogsRepositoryImpl) Update(item *models.GuildResourceTransactionLogs) error {
    return r.db.Save(item).Error
}

func (r *guildResourceTransactionLogsRepositoryImpl) BulkUpdate(items ...*models.GuildResourceTransactionLogs) error {
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

func (r *guildResourceTransactionLogsRepositoryImpl) Delete(id string) error {
    return r.db.Delete(&models.GuildResourceTransactionLogs{}, id).Error
}
