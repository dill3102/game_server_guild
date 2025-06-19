package repository

import (
    "gorm.io/gorm"
    models "github.com/dill3102/game_server_guild/internal/domain"
)

type UserCharactersRepository interface {
    GetByID(id string) (*models.UserCharacters, error)
    Create(item *models.UserCharacters) error
    Update(item *models.UserCharacters) error
    Delete(id string) error
}

type userCharactersRepositoryImpl struct {
    db *gorm.DB
}

func NewUserCharactersRepository(db *gorm.DB) UserCharactersRepository {
    return &userCharactersRepositoryImpl{db: db}
}

// GORM実装
func (r *userCharactersRepositoryImpl) GetByID(id string) (*models.UserCharacters, error) {
    var item models.UserCharacters
    if err := r.db.First(&item, id).Error; err != nil {
        return nil, err
    }
    return &item, nil
}

func (r *userCharactersRepositoryImpl) Create(item *models.UserCharacters) error {
    return r.db.Create(item).Error
}

func (r *userCharactersRepositoryImpl) BulkCreate(items ...*models.UserCharacters) error {
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

func (r *userCharactersRepositoryImpl) Update(item *models.UserCharacters) error {
    return r.db.Save(item).Error
}

func (r *userCharactersRepositoryImpl) BulkUpdate(items ...*models.UserCharacters) error {
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

func (r *userCharactersRepositoryImpl) Delete(id string) error {
    return r.db.Delete(&models.UserCharacters{}, id).Error
}
