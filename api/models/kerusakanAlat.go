package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type KerusakanAlat struct {
	gorm.Model
	Qty      int64  `gorm:"not null"              json:"qty"`
	Nama     string `gorm:"not null"             json:"nama"`
	Users_id int64  `gorm:"not null"              json:"users_id"`
	Status   string `gorm:"default:'Belum Diganti'" json:"status"`
}
type KerusakanAlatStruktur struct {
	Id        int64  `json:"id"`
	Qty       string `json:"qty"`
	Nama_user string `json:"nama_user"`
	Nama      string `json:"nama"`
	Status    string `json:"status"`
	CreatedAt time.Time
}

func (u *KerusakanAlat) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.ID == 0 {
			return errors.New("id is required")
		}
		if u.Nama == "" {
			return errors.New("Nama Alat is required")
		}
		if u.Qty == 0 {
			return errors.New("Qty is required")
		}
		return nil
	default:
		if u.Nama == "" {
			return errors.New("Nama Alat is required")
		}
		if u.Qty == 0 {
			return errors.New("Qty is required")
		}
		return nil
	}
}

func (u *KerusakanAlat) Save(db *gorm.DB) (*KerusakanAlat, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

var (
	selkerusakan    = "kerusakan_alats.*, users.name as nama_user"
	joinusertoeusak = "JOIN users on users.id = kerusakan_alats.users_id"
)

func (u *KerusakanAlat) FindInt(parameter string, condisi int, db *gorm.DB) (*[]KerusakanAlatStruktur, error) {
	data := []KerusakanAlatStruktur{}
	if err := db.Debug().Table("kerusakan_alats").Select(selkerusakan).Joins(joinusertoeusak).Where(parameter, condisi).Find(&data).Error; err != nil {
		return &[]KerusakanAlatStruktur{}, err
	}
	return &data, nil
}

func (u *KerusakanAlat) FindAll(parameter string, condisi int, db *gorm.DB) (*[]KerusakanAlatStruktur, error) {
	data := []KerusakanAlatStruktur{}
	if err := db.Debug().Table("kerusakan_alats").Select(selkerusakan).Joins(joinusertoeusak).Find(&data).Error; err != nil {
		return &[]KerusakanAlatStruktur{}, err
	}
	return &data, nil
}

func (v *KerusakanAlat) Delete(id string, db *gorm.DB) (bool, error) {
	if err := db.Unscoped().Table("kerusakan_alats").Where("id = ?", id).Delete(&KerusakanAlat{}).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (v *KerusakanAlat) UpdateDataKerusakan(db *gorm.DB) (*KerusakanAlat, error) {

	if err := db.Debug().Table("kerusakan_alats").Where("id = ?", v.ID).Updates(KerusakanAlat{
		Status: v.Status,
	}).Error; err != nil {
		return &KerusakanAlat{}, err
	}
	return v, nil
}
