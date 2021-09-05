package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type PeminjamanAlat struct {
	gorm.Model
	Qty      int64 `gorm:"not null"              json:"qty"`
	Alat_id  int64 `gorm:"not null"              json:"alat_id"`
	Users_id int64 `gorm:"not null"              json:"users_id"`
	Jenis    int64 `json:"jenis"`
}
type PeminjamanAlatStruktur struct {
	Id        int64  `json:"id"`
	Qty       string `json:"qty"`
	Nama_user string `json:"nama_user"`
	Nama_alat string `json:"nama_alat"`
	Alat_id   int64  `json:"alat_id"`
	Jenis     int64  `json:"jenis"`
	CreatedAt time.Time
}

func (u *PeminjamanAlat) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.ID == 0 {
			return errors.New("id is required")
		}
		if u.Alat_id == 0 {
			return errors.New("Alat is required")
		}
		if u.Qty == 0 {
			return errors.New("Qty is required")
		}
		return nil
	default:
		if u.Alat_id == 0 {
			return errors.New("Alat is required")
		}
		if u.Qty == 0 {
			return errors.New("Qty is required")
		}
		return nil
	}
}

func (u *PeminjamanAlat) Save(db *gorm.DB) (*PeminjamanAlat, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

var (
	selalat      = "peminjaman_alats.*, data_alats.name_alat as nama_alat, users.name as nama_user"
	joinalat     = "JOIN data_alats on data_alats.id = peminjaman_alats.alat_id"
	joinuserAlat = "JOIN users on users.id = peminjaman_alats.users_id"
)

func (u *PeminjamanAlat) FindInt(parameter string, condisi int, jenis string, db *gorm.DB) (*[]PeminjamanAlatStruktur, error) {
	data := []PeminjamanAlatStruktur{}
	if err := db.Debug().Table("peminjaman_alats").Select(selalat).Joins(joinalat).Joins(joinuserAlat).Where(parameter, condisi).Where("jenis = ?", jenis).Find(&data).Error; err != nil {
		return &[]PeminjamanAlatStruktur{}, err
	}
	return &data, nil
}

func (v *PeminjamanAlat) Delete(id string, db *gorm.DB) (bool, error) {
	if err := db.Unscoped().Table("peminjaman_alats").Where("id = ?", id).Delete(&PeminjamanAlat{}).Error; err != nil {
		return false, err
	}
	return true, nil
}
