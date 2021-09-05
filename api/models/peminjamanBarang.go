package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type PeminjamanBarang struct {
	gorm.Model
	Qty      int64 `gorm:"not null"              json:"qty"`
	Media_id int64 `gorm:"not null"              json:"media_id"`
	Users_id int64 `gorm:"not null"              json:"users_id"`
	Jenis    int64 `json:"jenis"`
	Status   int64 `gorm:"default:0"             json:"status"`
}
type PeminjamanBarangStruktur struct {
	Id         int64  `json:"id"`
	Qty        string `json:"qty"`
	Nama_user  string `json:"nama_user"`
	Nama_media string `json:"nama_media"`
	Media_id   int64  `json:"media_id"`
	Jenis      int64  `json:"jenis"`
	Status     int64  `json:"status"`
	CreatedAt  time.Time
}

func (u *PeminjamanBarang) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.ID == 0 {
			return errors.New("id is required")
		}
		if u.Media_id == 0 {
			return errors.New("Media is required")
		}
		if u.Qty == 0 {
			return errors.New("Qty is required")
		}
		return nil
	case "status":
		if u.Status < 0 {
			return errors.New("Status is required")
		}
		return nil
	default:
		if u.Media_id == 0 {
			return errors.New("Media is required")
		}
		if u.Qty == 0 {
			return errors.New("Qty is required")
		}
		return nil
	}
}

func (u *PeminjamanBarang) Save(db *gorm.DB) (*PeminjamanBarang, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

var (
	sel       = "peminjaman_barangs.*, data_media.name_media as nama_media, users.name as nama_user"
	joinmedia = "JOIN data_media on data_media.id = peminjaman_barangs.media_id"
	joinuser  = "JOIN users on users.id = peminjaman_barangs.users_id"
)

func (u *PeminjamanBarang) FindInt(parameter string, condisi int, jenis string, db *gorm.DB) (*[]PeminjamanBarangStruktur, error) {
	data := []PeminjamanBarangStruktur{}
	if err := db.Debug().Table("peminjaman_barangs").Select(sel).Joins(joinmedia).Joins(joinuser).Where(parameter, condisi).Where("jenis = ?", jenis).Find(&data).Error; err != nil {
		return &[]PeminjamanBarangStruktur{}, err
	}
	return &data, nil
}
func (u *PeminjamanBarang) FindStatus(parameter string, condisi int, jenis string, status string, db *gorm.DB) (*[]PeminjamanBarangStruktur, error) {
	data := []PeminjamanBarangStruktur{}
	if err := db.Debug().Table("peminjaman_barangs").Select(sel).Joins(joinmedia).Joins(joinuser).Where(parameter, condisi).Where("jenis = ?", jenis).Where("peminjaman_barangs.status = ?", status).Find(&data).Error; err != nil {
		return &[]PeminjamanBarangStruktur{}, err
	}
	return &data, nil
}

func (v *PeminjamanBarang) Delete(id string, db *gorm.DB) (bool, error) {
	if err := db.Unscoped().Table("peminjaman_barangs").Where("id = ?", id).Delete(&PeminjamanBarang{}).Error; err != nil {
		return false, err
	}
	return true, nil
}
func (v *PeminjamanBarang) Update(id uint, jenis int64, db *gorm.DB) (*PeminjamanBarang, error) {

	if err := db.Debug().Table("peminjaman_barangs").Where("users_id = ?", id).Where("jenis = ?", jenis).Updates(PeminjamanBarang{
		Status: v.Status,
	}).Error; err != nil {
		return &PeminjamanBarang{}, err
	}
	return v, nil
}
