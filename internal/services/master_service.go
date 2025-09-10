package services

import (
	"gorm.io/gorm"
)

type MasterService struct { db *gorm.DB }
func NewMasterService(db *gorm.DB) *MasterService { return &MasterService{db} }
func (s *MasterService) DB() *gorm.DB { return s.db }
