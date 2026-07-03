package repositories

import (
	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	DB *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{
		DB: db,
	}
}

func (r *BaseRepository[T]) Create(entity *T) error {
	return r.DB.Create(entity).Error
}

func (r *BaseRepository[T]) GetAll() ([]T, error) {
	var entities []T
	err := r.DB.Find(&entities).Error
	return entities, err
}

func (r *BaseRepository[T]) GetByID(id string) (T, error) {
	var entity T
	err := r.DB.First(&entity, id).Error
	return entity, err
}

func (r *BaseRepository[T]) FindWhere(query interface{}, args ...interface{}) ([]T, error) {
	var entities []T
	err := r.DB.Where(query, args...).Find(&entities).Error
	return entities, err
}

func (r *BaseRepository[T]) Update(id string, entity *T) error {
	var model T
	return r.DB.Model(&model).Where("id = ?", id).Updates(entity).Error
}

func (r *BaseRepository[T]) Delete(id string) error {
	var entity T
	return r.DB.Delete(&entity, id).Error
}
