package models

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
)

type DataPenelitian struct {
	gorm.Model
	Data_alats_id int64  `gorm:"size:100;not null" json:"data_alats_id"`
	Nama_alat     string `gorm:"size:100;not null" json:"nama_alat"`
	Qty           int64  `gorm:"not null" json:"qty"`
	Tanggal       string `gorm:"size:100;not null" json:"tanggal"`
	Jenis         string `gorm:"" json:"jenis"`
	Status        string `gorm:"" json:"status"`
	Created_by    int64  `gorm:"not null" json:"created_by"`
	Updated_by    int64  `gorm:"" json:"updated_by"`
}
type StukturDataPenelitian struct {
	Id            int64  `json:"id"`
	Data_alats_id int64  `json:"data_alats_id"`
	Nama_alat     string `json:"nama_alat"`
	Qty           string `json:"qty"`
	Tanggal       string `json:"tanggal"`
	Jenis         string `json:"jenis"`
	Status        string `json:"status"`
	Created_by    int64  `json:"created_by"`
	Updated_by    int64  `json:"updated_by"`
}

type StrukturDataShow struct {
	Rusak       interface{}
	Tidak_rusak interface{}
}

func (u *DataPenelitian) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.ID == 0 {
			return errors.New("Data Media is required")
		}
		if u.Status == "" {
			return errors.New("Status is required")
		}
		return nil
	default:
		if u.Data_alats_id <= 0 {
			return errors.New("Data Alat is required")
		}
		if u.Qty <= 0 {
			return errors.New("Merek media is required")
		}

		if u.Tanggal == "" {
			return errors.New("Tanggal is required")
		}
		if u.Jenis == "" {
			return errors.New("Jenis is required")
		}
		if u.Status == "" {
			return errors.New("Status is required")
		}
		return nil
	}
}

func (u *DataPenelitian) SaveDataPenelitian(db *gorm.DB) (*DataPenelitian, error) {
	var err error

	err = db.Debug().Create(&u).Error
	if err != nil {
		return &DataPenelitian{}, err
	}
	return u, nil
}

func (u *DataPenelitian) GetDataPenelitian(db *gorm.DB, parameter string, data string) (*DataPenelitian, error) {
	account := &DataPenelitian{}
	if err := db.Debug().Table("data_penelitians").Where(parameter, data).First(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}
func (u *DataPenelitian) GetAll(value string, parameter string, db *gorm.DB) (*[]DataPenelitian, error) {
	account := []DataPenelitian{}
	var condition = value + " LIKE ?"
	if err := db.Debug().Table("data_penelitians").Where(condition, "%"+parameter+"%").Find(&account).Error; err != nil {
		return nil, errors.New("Data Tidak di temukan")
	}
	return &account, nil
}
func (u *DataPenelitian) Getint(value string, parameter int, db *gorm.DB) (*StrukturDataShow, error) {
	account := []DataPenelitian{}
	var condition = value + " = ?"
	if err := db.Debug().Table("data_penelitians").Where(condition, parameter).Where("jenis = ?", "Rusak").Find(&account).Error; err != nil {
		return nil, errors.New("Data Tidak di temukan")
	}
	lol := []DataPenelitian{}
	if err := db.Debug().Table("data_penelitians").Where(condition, parameter).Where("jenis = ?", "Tidak Rusak").Find(&lol).Error; err != nil {
		return nil, errors.New("Data Tidak di temukan")
	}
	data := StrukturDataShow{

		Rusak:       account,
		Tidak_rusak: lol,
	}
	return &data, nil
}

func (v *DataPenelitian) UpdateStatusDataPenelitian(id uint, db *gorm.DB) (*DataPenelitian, error) {

	if err := db.Debug().Table("data_penelitians").Where("id = ?", id).Updates(DataPenelitian{
		Status:     v.Status,
		Updated_by: v.Updated_by,
	}).Error; err != nil {
		return &DataPenelitian{}, err
	}
	return v, nil
}

func (v *DataPenelitian) Delete(id string, db *gorm.DB) (bool, error) {
	if err := db.Unscoped().Table("data_penelitians").Where("id = ?", id).Delete(&DataPenelitian{}).Error; err != nil {
		return false, err
	}
	return true, nil
}
