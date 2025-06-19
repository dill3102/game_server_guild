package repository

import (
    "gorm.io/gorm"
    models "github.com/dill3102/game_server_guild/internal/domain"
)

type UserFacilityMappingRepository interface {
    GetByID(id string) (*models.UserFacilityMapping, error)
    Create(item *models.UserFacilityMapping) error
    Update(item *models.UserFacilityMapping) error
    Delete(id string) error
}

type userFacilityMappingRepositoryImpl struct {
    db *gorm.DB
}

func NewUserFacilityMappingRepository(db *gorm.DB) UserFacilityMappingRepository {
    return &userFacilityMappingRepositoryImpl{db: db}
}

// GORM実装
func (r *userFacilityMappingRepositoryImpl) GetByID(id string) (*models.UserFacilityMapping, error) {
    var item models.UserFacilityMapping
    if err := r.db.First(&item, id).Error; err != nil {
        return nil, err
    }
    return &item, nil
}

func (r *userFacilityMappingRepositoryImpl) Create(item *models.UserFacilityMapping) error {
    return r.db.Create(item).Error
}

func (r *userFacilityMappingRepositoryImpl) BulkCreate(items ...*models.UserFacilityMapping) error {
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

func (r *userFacilityMappingRepositoryImpl) Update(item *models.UserFacilityMapping) error {
    return r.db.Save(item).Error
}

func (r *userFacilityMappingRepositoryImpl) BulkUpdate(items ...*models.UserFacilityMapping) error {
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

func (r *userFacilityMappingRepositoryImpl) Delete(id string) error {
    return r.db.Delete(&models.UserFacilityMapping{}, id).Error
}
