package models

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
)

type DataMedia struct {
	gorm.Model
	Name_media    string `gorm:"size:100;not null"              json:"name_media"`
	Satuan        string `gorm:"size:100;not null" 				json:"satuan"`
	Merek_media   string `gorm:"size:100;not null"              json:"merek_media"`
	Distributor   string `gorm:"size:100;not null"       		json:"distributor"`
	Tanggal_masuk string `gorm:"size:100;not null"          	json:"tanggal_masuk"`
	Expired       string `gorm:"size:1000;not null"          	json:"expired"`
	Status        string `gorm:"size:1000;not null"          	json:"status"`
	Created_by    int64  `gorm:"size:1000;not null"          	json:"created_by"`
	Updated_by    int64  `gorm:"size:1000;"          			json:"updated_by"`
}
type StukturDataMedia struct {
	Id            int64  `json:"id"`
	Name_media    string `gorm:"size:100;not null"              json:"name_media"`
	Satuan        string `gorm:"size:100;not null" 				json:"satuan"`
	Merek_media   string `gorm:"size:100;not null"              json:"merek_media"`
	Distributor   string `gorm:"size:100;not null"       		json:"distributor"`
	Tanggal_masuk string `gorm:"size:100;not null"          	json:"tanggal_masuk"`
	Expired       string `gorm:"size:1000;not null"          	json:"expired"`
	Created_by    int64  `gorm:"size:1000;not null"          	json:"created_by"`
	Updated_by    int64  `gorm:"size:1000;"          			json:"updated_by"`
}

func (u *DataMedia) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.ID == 0 {
			return errors.New("Data Media is required")
		}
		if u.Name_media == "" {
			return errors.New("Name media is required")
		}
		if u.Satuan == "" {
			return errors.New("satuan is required")
		}
		if u.Merek_media == "" {
			return errors.New("Merek media is required")
		}
		if u.Distributor == "" {
			return errors.New("Distributor is required")
		}
		if u.Tanggal_masuk == "" {
			return errors.New("Tanggal masuk is required")
		}
		if u.Expired == "" {
			return errors.New("Expired is required")
		}
		if u.Status == "" {
			return errors.New("Status is required")
		}
		return nil
	default:
		if u.Name_media == "" {
			return errors.New("Name media is required")
		}
		if u.Satuan == "" {
			return errors.New("satuan is required")
		}
		if u.Merek_media == "" {
			return errors.New("Merek media is required")
		}
		if u.Distributor == "" {
			return errors.New("Distributor is required")
		}
		if u.Tanggal_masuk == "" {
			return errors.New("Tanggal masuk is required")
		}
		if u.Expired == "" {
			return errors.New("Expired is required")
		}
		if u.Status == "" {
			return errors.New("Status is required")
		}
		return nil
	}
}

func (u *DataMedia) SaveDataMedia(db *gorm.DB) (*DataMedia, error) {
	var err error

	err = db.Debug().Create(&u).Error
	if err != nil {
		return &DataMedia{}, err
	}
	return u, nil
}

func (u *DataMedia) GetDataMedia(db *gorm.DB, parameter string, data string) (*DataMedia, error) {
	account := &DataMedia{}
	if err := db.Debug().Table("data_media").Where(parameter, data).First(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}
func (u *DataMedia) GetAll(value string, parameter string, db *gorm.DB) (*[]DataMedia, error) {
	account := []DataMedia{}
	var condition = value + " LIKE ?"
	if err := db.Debug().Table("data_media").Where(condition, "%"+parameter+"%").Find(&account).Error; err != nil {
		return nil, errors.New("Data Tidak di temukan")
	}
	return &account, nil
}

func (v *DataMedia) UpdateDataMedia(id uint, db *gorm.DB) (*DataMedia, error) {

	if err := db.Debug().Table("data_media").Where("id = ?", id).Updates(DataMedia{
		Name_media:    v.Name_media,
		Satuan:        v.Satuan,
		Merek_media:   v.Merek_media,
		Distributor:   v.Distributor,
		Tanggal_masuk: v.Tanggal_masuk,
		Expired:       v.Expired,
		Updated_by:    v.Updated_by,
		Status:        v.Status,
	}).Error; err != nil {
		return &DataMedia{}, err
	}
	return v, nil
}

func (v *DataMedia) Delete(id string, db *gorm.DB) (bool, error) {
	if err := db.Unscoped().Table("data_media").Where("id = ?", id).Delete(&DataMedia{}).Error; err != nil {
		return false, err
	}
	return true, nil
}
