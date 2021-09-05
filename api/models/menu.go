package models

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
)

type Menu struct {
	gorm.Model
	Nama       string `gorm:"sizw:100:not null" json:"nama"`
	Role       string `gorm:"size:100:not null" json:"role"`
	Navigation string `gorm:"size:100;not null;unique_index"              json:"navigation"`
}
type StukturMenu struct {
	Id         int64  `json:"id"`
	Role       string `json:"role"`
	Nama       string `json:"nama"`
	Navigation string `json:"navigation"`
}

func (u *Menu) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.ID == 0 {
			return errors.New("id is required")
		}
		if u.Role == "" {
			return errors.New("Role is required")
		}
		if u.Nama == "" {
			return errors.New("nama is required")
		}
		if u.Navigation == "" {
			return errors.New("navigation is required")
		}
		switch u.Role {
		case "Ka Lab":
			return nil
		case "Laboran":
			return nil
		case "Mahasiswa":
			return nil
		}
		return errors.New("role type not found ")
	default:
		if u.Role == "" {
			return errors.New("Role is required")
		}
		if u.Nama == "" {
			return errors.New("nama is required")
		}
		if u.Navigation == "" {
			return errors.New("navigation is required")
		}
		switch u.Role {
		case "Ka Lab":
			return nil
		case "Laboran":
			return nil
		case "Mahasiswa":
			return nil
		}
		return errors.New("role type not found ")
	}
}

func (u *Menu) Save(db *gorm.DB) (*Menu, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Menu{}, err
	}
	return u, nil
}
func (u *Menu) All(db *gorm.DB) (*[]Menu, error) {
	data := []Menu{}
	if err := db.Debug().Table("menus").Find(&data).Error; err != nil {
		return &[]Menu{}, err
	}
	return &data, nil
}
func (u *Menu) FindString(parameter string, condisi string, db *gorm.DB) (*[]Menu, error) {
	data := []Menu{}
	if err := db.Debug().Table("menus").Where(parameter, condisi).Find(&data).Error; err != nil {
		return &[]Menu{}, err
	}
	return &data, nil
}
func (u *Menu) FindInt(parameter string, condisi int, db *gorm.DB) (*[]Menu, error) {
	data := []Menu{}
	if err := db.Debug().Table("menus").Where(parameter, condisi).Find(&data).Error; err != nil {
		return &[]Menu{}, err
	}
	return &data, nil
}

func (v *Menu) Update(db *gorm.DB) (*Menu, error) {

	if err := db.Debug().Table("menus").Where("id = ?", v.ID).Updates(Menu{
		Role:       v.Role,
		Nama:       v.Nama,
		Navigation: v.Navigation,
	}).Error; err != nil {
		return &Menu{}, err
	}
	return v, nil
}
