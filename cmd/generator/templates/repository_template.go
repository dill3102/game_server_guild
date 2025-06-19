package repository

import (
    "gorm.io/gorm"
    models "github.com/dill3102/game_server_guild/internal/domain"
)

type {ModelName}Repository interface {
    GetByID(id string) (*models.{ModelName}, error)
    Create(item *models.{ModelName}) error
    Update(item *models.{ModelName}) error
    Delete(id string) error
}

type {ImplName}RepositoryImpl struct {
    db *gorm.DB
}

func New{ModelName}Repository(db *gorm.DB) {ModelName}Repository {
    return &{ImplName}RepositoryImpl{db: db}
}

// GORM実装
func (r *{ImplName}RepositoryImpl) GetByID(id string) (*models.{ModelName}, error) {
    var item models.{ModelName}
    if err := r.db.First(&item, id).Error; err != nil {
        return nil, err
    }
    return &item, nil
}

func (r *{ImplName}RepositoryImpl) Create(item *models.{ModelName}) error {
    return r.db.Create(item).Error
}

func (r *{ImplName}RepositoryImpl) BulkCreate(items ...*models.{ModelName}) error {
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

func (r *{ImplName}RepositoryImpl) Update(item *models.{ModelName}) error {
    return r.db.Save(item).Error
}

func (r *{ImplName}RepositoryImpl) BulkUpdate(items ...*models.{ModelName}) error {
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

func (r *{ImplName}RepositoryImpl) Delete(id string) error {
    return r.db.Delete(&models.{ModelName}{}, id).Error
}
