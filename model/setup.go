package model

import (
	"errors"

	"github.com/Isiah998/new-api/common"
	"gorm.io/gorm"
)

type Setup struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	Version       string `json:"version" gorm:"type:varchar(50);not null"`
	InitializedAt int64  `json:"initialized_at" gorm:"type:bigint;not null"`
}

func GetSetup() *Setup {
	var setup Setup
	err := DB.First(&setup).Error
	if err != nil {
		return nil
	}
	return &setup
}

// InitializeSystem claims setup with a fixed primary key and commits the root
// account plus initial options in one transaction. The unique claim prevents
// two application instances from both completing first-run initialization.
func InitializeSystem(setup Setup, rootUser *User, options map[string]string) error {
	setup.ID = 1
	err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&setup).Error; err != nil {
			return err
		}

		var rootCount int64
		if err := tx.Model(&User{}).Where("role = ?", common.RoleRootUser).Count(&rootCount).Error; err != nil {
			return err
		}
		if rootCount == 0 {
			if rootUser == nil {
				return errors.New("root user is required")
			}
			if err := tx.Create(rootUser).Error; err != nil {
				return err
			}
		}

		for key, value := range options {
			option := Option{Key: key}
			if err := tx.FirstOrCreate(&option, Option{Key: key}).Error; err != nil {
				return err
			}
			option.Value = value
			if err := tx.Save(&option).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		if GetSetup() != nil {
			return ErrSetupCompleted
		}
		return err
	}

	for key, value := range options {
		if err := updateOptionMap(key, value); err != nil {
			common.SysLog("failed to apply initialized option " + key + ": " + err.Error())
		}
	}
	return nil
}
