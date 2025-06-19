package repository

import (
    "gorm.io/gorm"
    "your/module/models"
)

type UserResourcesRepository interface {
    GetByID(id string) (*models.UserResources, error)
    Create(item *models.UserResources) error
    Update(item *models.UserResources) error
    Delete(id string) error
}

type userResourcesRepositoryImpl struct {
    db *gorm.DB
}

func NewUserResourcesRepository(db *gorm.DB) UserResourcesRepository {
    return &userResourcesRepositoryImpl{db: db}
}

// GORM実装
func (r *userResourcesRepositoryImpl) GetByID(id string) (*models.UserResources, error) {
    var item models.UserResources
    if err := r.db.First(&item, id).Error; err != nil {
        return nil, err
    }
    return &item, nil
}

func (r *userResourcesRepositoryImpl) Create(item *models.UserResources) error {
    return r.db.Create(item).Error
}

func (r *userResourcesRepositoryImpl) BulkCreate(items ...*models.UserResources) error {
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

func (r *userResourcesRepositoryImpl) Update(item *models.UserResources) error {
    return r.db.Save(item).Error
}

func (r *userResourcesRepositoryImpl) BulkUpdate(items ...*models.UserResources) error {
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

func (r *userResourcesRepositoryImpl) Delete(id string) error {
    return r.db.Delete(&models.UserResources{}, id).Error
}
