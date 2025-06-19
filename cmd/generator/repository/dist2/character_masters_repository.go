package repository

import (
    "gorm.io/gorm"
    "your/module/models"
)

type CharacterMastersRepository interface {
    GetByID(id string) (*models.CharacterMasters, error)
    Create(item *models.CharacterMasters) error
    Update(item *models.CharacterMasters) error
    Delete(id string) error
}

type characterMastersRepositoryImpl struct {
    db *gorm.DB
}

func NewCharacterMastersRepository(db *gorm.DB) CharacterMastersRepository {
    return &characterMastersRepositoryImpl{db: db}
}

// GORM実装
func (r *characterMastersRepositoryImpl) GetByID(id string) (*models.CharacterMasters, error) {
    var item models.CharacterMasters
    if err := r.db.First(&item, id).Error; err != nil {
        return nil, err
    }
    return &item, nil
}

func (r *characterMastersRepositoryImpl) Create(item *models.CharacterMasters) error {
    return r.db.Create(item).Error
}

func (r *characterMastersRepositoryImpl) BulkCreate(items ...*models.CharacterMasters) error {
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

func (r *characterMastersRepositoryImpl) Update(item *models.CharacterMasters) error {
    return r.db.Save(item).Error
}

func (r *characterMastersRepositoryImpl) BulkUpdate(items ...*models.CharacterMasters) error {
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

func (r *characterMastersRepositoryImpl) Delete(id string) error {
    return r.db.Delete(&models.CharacterMasters{}, id).Error
}
