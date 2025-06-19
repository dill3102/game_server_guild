package repository

import (
    "gorm.io/gorm"
    models "github.com/dill3102/game_server_guild/internal/domain"
)

type UpgradeCostMastersRepository interface {
    GetByID(id string) (*models.UpgradeCostMasters, error)
    Create(item *models.UpgradeCostMasters) error
    Update(item *models.UpgradeCostMasters) error
    Delete(id string) error
}

type upgradeCostMastersRepositoryImpl struct {
    db *gorm.DB
}

func NewUpgradeCostMastersRepository(db *gorm.DB) UpgradeCostMastersRepository {
    return &upgradeCostMastersRepositoryImpl{db: db}
}

// GORM実装
func (r *upgradeCostMastersRepositoryImpl) GetByID(id string) (*models.UpgradeCostMasters, error) {
    var item models.UpgradeCostMasters
    if err := r.db.First(&item, id).Error; err != nil {
        return nil, err
    }
    return &item, nil
}

func (r *upgradeCostMastersRepositoryImpl) Create(item *models.UpgradeCostMasters) error {
    return r.db.Create(item).Error
}

func (r *upgradeCostMastersRepositoryImpl) BulkCreate(items ...*models.UpgradeCostMasters) error {
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

func (r *upgradeCostMastersRepositoryImpl) Update(item *models.UpgradeCostMasters) error {
    return r.db.Save(item).Error
}

func (r *upgradeCostMastersRepositoryImpl) BulkUpdate(items ...*models.UpgradeCostMasters) error {
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

func (r *upgradeCostMastersRepositoryImpl) Delete(id string) error {
    return r.db.Delete(&models.UpgradeCostMasters{}, id).Error
}
