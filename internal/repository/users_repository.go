package repository

import (
    "gorm.io/gorm"
    models "github.com/dill3102/game_server_guild/internal/domain"
)

type UsersRepository interface {
    GetByID(id string) (*models.Users, error)
    Create(item *models.Users) error
    Update(item *models.Users) error
    Delete(id string) error
}

type usersRepositoryImpl struct {
    db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
    return &usersRepositoryImpl{db: db}
}

// GORM実装
func (r *usersRepositoryImpl) GetByID(id string) (*models.Users, error) {
    var item models.Users
    if err := r.db.First(&item, id).Error; err != nil {
        return nil, err
    }
    return &item, nil
}

func (r *usersRepositoryImpl) Create(item *models.Users) error {
    return r.db.Create(item).Error
}

func (r *usersRepositoryImpl) BulkCreate(items ...*models.Users) error {
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

func (r *usersRepositoryImpl) Update(item *models.Users) error {
    return r.db.Save(item).Error
}

func (r *usersRepositoryImpl) BulkUpdate(items ...*models.Users) error {
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

func (r *usersRepositoryImpl) Delete(id string) error {
    return r.db.Delete(&models.Users{}, id).Error
}
