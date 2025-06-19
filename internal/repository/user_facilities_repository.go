package repository

import (
    "gorm.io/gorm"
    models "github.com/dill3102/game_server_guild/internal/domain"
)

type UserFacilitiesRepository interface {
    GetByID(id string) (*models.UserFacilities, error)
    Create(item *models.UserFacilities) error
    Update(item *models.UserFacilities) error
    Delete(id string) error
}

type userFacilitiesRepositoryImpl struct {
    db *gorm.DB
}

func NewUserFacilitiesRepository(db *gorm.DB) UserFacilitiesRepository {
    return &userFacilitiesRepositoryImpl{db: db}
}

// GORM実装
func (r *userFacilitiesRepositoryImpl) GetByID(id string) (*models.UserFacilities, error) {
    var item models.UserFacilities
    if err := r.db.First(&item, id).Error; err != nil {
        return nil, err
    }
    return &item, nil
}

func (r *userFacilitiesRepositoryImpl) Create(item *models.UserFacilities) error {
    return r.db.Create(item).Error
}

func (r *userFacilitiesRepositoryImpl) BulkCreate(items ...*models.UserFacilities) error {
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

func (r *userFacilitiesRepositoryImpl) Update(item *models.UserFacilities) error {
    return r.db.Save(item).Error
}

func (r *userFacilitiesRepositoryImpl) BulkUpdate(items ...*models.UserFacilities) error {
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

func (r *userFacilitiesRepositoryImpl) Delete(id string) error {
    return r.db.Delete(&models.UserFacilities{}, id).Error
}
