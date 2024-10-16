package link

import (
	"context"

	"gorm.io/gorm"
)

type LinkRepository struct {
	DB *gorm.DB
}

func NewLinkRepository(db *gorm.DB) *LinkRepository {
	return &LinkRepository{db}
}

func (r *LinkRepository) Create(ctx context.Context, id string, entity *Link) error {
	tx := r.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Save(entity).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *LinkRepository) Delete(ctx context.Context, id string, data string) error {
	tx := r.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	link := new(Link)
	if err := r.DB.Where("user_id = ? AND short_link = ?", id, data).First(link).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(link).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *LinkRepository) FindByLink(ctx context.Context, entity *Link, link string) error {
	tx := r.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Where("short_link = ?", link).Take(entity).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *LinkRepository) ListByUUID(ctx context.Context, id string) ([]Link, error) {
	var Links []Link
	tx := r.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	if err := tx.Where("user_id = ?", id).Take(&Links).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return Links, nil
}
