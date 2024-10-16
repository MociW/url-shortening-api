package user

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(ctx context.Context, entity *User) error {
	tx := r.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	var total int64
	if err := r.DB.Model(entity).Where("email = ?", entity.Email).Count(&total).Error; err != nil {
		return err
	}

	if total == 0 {
		if err := tx.Save(entity).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		tx.Rollback()
		return fmt.Errorf("email already exists")
	}

	return tx.Commit().Error
}

func (r *UserRepository) Delete(ctx context.Context, email string, password string) error {
	tx := r.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	user := new(User)
	if err := r.DB.Where("email = ? AND password = ?", email, password).First(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *UserRepository) Update(ctx context.Context, entity *User, uuid any) error {
	tx := r.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	user := new(User)
	if err := r.DB.Where("user_id = ?", uuid).First(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	user.Username = entity.Username
	user.Email = entity.Email
	user.Password = entity.Password

	if err := tx.Save(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *UserRepository) FindByEmail(ctx context.Context, entity *User, email any) error {
	tx := r.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Where("email = ?", email).First(entity).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
