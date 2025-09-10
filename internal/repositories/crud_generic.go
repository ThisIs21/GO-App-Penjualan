package repositories

import (
	"gorm.io/gorm"
)

type CRUD[T any] struct{ db *gorm.DB }

func NewCRUD[T any](db *gorm.DB) *CRUD[T] { return &CRUD[T]{db: db} }

func (r *CRUD[T]) List(dst *[]T) error  { return r.db.Find(dst).Error }
func (r *CRUD[T]) Get(id uint, dst *T) error { return r.db.First(dst, id).Error }
func (r *CRUD[T]) Create(m *T) error    { return r.db.Create(m).Error }
func (r *CRUD[T]) Update(id uint, m *T) error { return r.db.Model(m).Where("id = ?", id).Updates(m).Error }
func (r *CRUD[T]) Delete(id uint, m *T) error { return r.db.Delete(m, id).Error }
