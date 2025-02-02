// The Base Repository is a generic repository for CRUD operations
// That all other repositories will inherit from. Avoiding code duplication
package repositories

import (
	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{db: db};
}

func (r *BaseRepository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error;
}

func (r *BaseRepository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error;
}

func (r *BaseRepository[T]) Delete(entity *T) error {
	return r.db.Delete(entity).Error;
}

func (r *BaseRepository[T]) FindAll(entities *[]T) error {
	return r.db.Find(entities).Error;
}

func (r *BaseRepository[T]) FindById(id int64, entity *T) error {
	return r.db.First(entity, id).Error;
}