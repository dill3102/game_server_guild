package repository

import (
    "gorm.io/gorm"
    "your/module/models"
)

type FacilityMastersRepository interface {
    GetByID(id string) (*models.FacilityMasters, error)
    Create(item *models.FacilityMasters) error
    Update(item *models.FacilityMasters) error
    Delete(id string) error
}

type facilityMastersRepositoryImpl struct {
    db *gorm.DB
}

func NewFacilityMastersRepository(db *gorm.DB) FacilityMastersRepository {
    return &facilityMastersRepositoryImpl{db: db}
}

// GORM実装
func (r *facilityMastersRepositoryImpl) GetByID(id string) (*models.FacilityMasters, error) {
    var item models.FacilityMasters
    if err := r.db.First(&item, id).Error; err != nil {
        return nil, err
    }
    return &item, nil
}

func (r *facilityMastersRepositoryImpl) Create(item *models.FacilityMasters) error {
    return r.db.Create(item).Error
}

func (r *facilityMastersRepositoryImpl) BulkCreate(items ...*models.FacilityMasters) error {
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

func (r *facilityMastersRepositoryImpl) Update(item *models.FacilityMasters) error {
    return r.db.Save(item).Error
}

func (r *facilityMastersRepositoryImpl) BulkUpdate(items ...*models.FacilityMasters) error {
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

func (r *facilityMastersRepositoryImpl) Delete(id string) error {
    return r.db.Delete(&models.FacilityMasters{}, id).Error
}
