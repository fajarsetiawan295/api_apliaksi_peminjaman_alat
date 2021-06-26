package models

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
)

type DataAlat struct {
	gorm.Model
	Name_alat     string `gorm:"size:100;not null"              json:"name_alat"`
	Kapasitas     string `gorm:"size:100;not null" 				json:"kapasitas"`
	Merek_alat    string `gorm:"size:100;not null"              json:"merek_alat"`
	Distributor   string `gorm:"size:100;not null"       		json:"distributor"`
	Tanggal_masuk string `gorm:"size:100;not null"          	json:"tanggal_masuk"`
	Qty           int64  `gorm:"size:1000;not null"          	json:"qty"`
	Created_by    int64  `gorm:"size:1000;not null"          	json:"created_by"`
	Updated_by    int64  `gorm:"size:1000;"          			json:"updated_by"`
}
type StukturDataAlat struct {
	Id            int64  `json:"id"`
	Name_alat     string `json:"name_alat"`
	Kapasitas     string `json:"kapasitas"`
	Merek_alat    string `json:"merek_alat"`
	Distributor   string `json:"distributor"`
	Tanggal_masuk string `json:"tanggal_masuk"`
	Qty           int64  `json:"qty"`
	Created_by    int64  `json:"created_by"`
	Updated_by    int64  `json:"updated_by"`
}

func (u *DataAlat) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.ID == 0 {
			return errors.New("Data Media is required")
		}
		if u.Name_alat == "" {
			return errors.New("Name media is required")
		}
		if u.Kapasitas == "" {
			return errors.New("Kapasitas is required")
		}
		if u.Merek_alat == "" {
			return errors.New("Merek media is required")
		}
		if u.Distributor == "" {
			return errors.New("Distributor is required")
		}
		if u.Tanggal_masuk == "" {
			return errors.New("Tanggal masuk is required")
		}

		if u.Qty <= 0 {
			return errors.New("Qty is required")
		}
		return nil
	default:
		if u.Name_alat == "" {
			return errors.New("Name media is required")
		}
		if u.Kapasitas == "" {
			return errors.New("Kapasitas is required")
		}
		if u.Merek_alat == "" {
			return errors.New("Merek media is required")
		}
		if u.Distributor == "" {
			return errors.New("Distributor is required")
		}
		if u.Tanggal_masuk == "" {
			return errors.New("Tanggal masuk is required")
		}

		if u.Qty <= 0 {
			return errors.New("Qty is required")
		}
		return nil
	}
}

func (u *DataAlat) SaveDataAlat(db *gorm.DB) (*DataAlat, error) {
	var err error

	err = db.Debug().Create(&u).Error
	if err != nil {
		return &DataAlat{}, err
	}
	return u, nil
}

func (u *DataAlat) GetDataAlat(db *gorm.DB, parameter string, data string) (*DataAlat, error) {
	account := &DataAlat{}
	if err := db.Debug().Table("data_alats").Where(parameter, data).First(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}
func (u *DataAlat) GetDataAlatInt(db *gorm.DB, parameter string, data int) (*DataAlat, error) {
	account := &DataAlat{}
	if err := db.Debug().Table("data_alats").Where(parameter, data).First(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}
func (u *DataAlat) GetAll(value string, parameter string, db *gorm.DB) (*[]DataAlat, error) {
	account := []DataAlat{}
	var condition = value + " LIKE ?"
	if err := db.Debug().Table("data_alats").Where(condition, "%"+parameter+"%").Find(&account).Error; err != nil {
		return nil, errors.New("Data Tidak di temukan")
	}
	return &account, nil
}

func (v *DataAlat) UpdateDataAlat(id uint, db *gorm.DB) (*DataAlat, error) {

	if err := db.Debug().Table("data_alats").Where("id = ?", id).Updates(DataAlat{
		Name_alat:     v.Name_alat,
		Kapasitas:     v.Kapasitas,
		Merek_alat:    v.Merek_alat,
		Distributor:   v.Distributor,
		Tanggal_masuk: v.Tanggal_masuk,
		Updated_by:    v.Updated_by,
		Qty:           v.Qty,
	}).Error; err != nil {
		return &DataAlat{}, err
	}
	return v, nil
}

func (v *DataAlat) Delete(id string, db *gorm.DB) (bool, error) {
	if err := db.Unscoped().Table("data_alats").Where("id = ?", id).Delete(&DataAlat{}).Error; err != nil {
		return false, err
	}
	return true, nil
}
