package models

import (
	"errors"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Email        string `gorm:"type:varchar(100);unique_index" 	json:"email"`
	Name         string `gorm:"size:100;not null"              	json:"name"`
	Nama_lengkap string `gorm:"size:100;"              	json:"nama_lengkap"`
	Nomor_hp     string `gorm:"size:100;not null;unique_index" 	json:"nomor_hp"`
	Password     string `gorm:"size:100;not null"              	json:"password"`
	Role         string `gorm:"size:100;not null"         		json:"role"`
	Npm          string `gorm:"size:100;not null"          		json:"npm"`
	Foto         string `gorm:"size:1000;not null"          	json:"foto"`
}
type StukturUser struct {
	Id           int64  `json:"id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Nama_lengkap string `json:"nama_lengkap"`
	Nomor_hp     string `json:"nomor_hp"`
	Password     string `json:"password"`
	Role         string `json:"role"`
	Npm          string `json:"npm"`
	Foto         string `json:"foto"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return errors.New("password incorrect")
	}
	return nil
}

func (u *User) BeforeSave() error {
	password := strings.TrimSpace(u.Password)
	hashedpassword, err := HashPassword(password)
	if err != nil {
		return err
	}
	u.Password = string(hashedpassword)
	return nil
}

func (u *User) Prepare() {
	u.Email = strings.TrimSpace(u.Email)
	u.Name = strings.TrimSpace(u.Name)
	u.Nomor_hp = strings.TrimSpace(u.Nomor_hp)
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "login":
		if u.Nomor_hp == "" && u.Email == "" {
			return errors.New("Nomor Handphone or Email is required")
		}
		if u.Password == "" {
			return errors.New("Password is required")
		}
		return nil
	default:
		if u.Name == "" {
			return errors.New("Name is required")
		}
		if u.Nomor_hp == "" {
			return errors.New("nomor handpone is required")
		}
		if u.Email == "" {
			return errors.New("Email is required")
		}
		if u.Password == "" {
			return errors.New("Password is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		if u.Role == "" {

			return errors.New("role is required")
		}
		if u.Npm == "" {

			return errors.New("NPM is required")
		}
		if u.Nama_lengkap == "" {

			return errors.New("Nama Lengkap is required")
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

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error

	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) GetUser(db *gorm.DB, parameter string, data string) (*User, error) {
	account := &User{}
	if err := db.Debug().Table("users").Where(parameter, data).First(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}

func GetAllUsers(db *gorm.DB) (*[]StukturUser, error) {
	users := []StukturUser{}
	if err := db.Debug().Table("users").Find(&users).Error; err != nil {
		return &[]StukturUser{}, err
	}
	return &users, nil
}
func GetAllbyRole(parameter string, db *gorm.DB) (*[]StukturUser, error) {
	users := []StukturUser{}
	if err := db.Debug().Table("users").Where("role = ?", parameter).Find(&users).Error; err != nil {
		return &[]StukturUser{}, err
	}
	return &users, nil
}

func Getfinduser(id uint, db *gorm.DB) (*StukturUser, error) {
	venue := &StukturUser{}
	if err := db.Debug().Table("users").Where("id = ?", id).First(venue).Error; err != nil {
		return nil, err
	}
	return venue, nil
}

func (v *User) UpdateUser(id uint, pa string, db *gorm.DB) (*User, error) {

	hashedpassword, _ := HashPassword(pa)

	if err := db.Debug().Table("users").Where("id = ?", id).Updates(User{
		Password: string(hashedpassword),
	}).Error; err != nil {
		return &User{}, err
	}
	return v, nil
}
func (v *User) UpdateForgotPassword(id string, pa string, db *gorm.DB) (*User, error) {

	hashedpassword, _ := HashPassword(pa)

	if err := db.Debug().Table("users").Where("nomor_hp = ?", id).Updates(User{
		Password: string(hashedpassword),
	}).Error; err != nil {
		return &User{}, err
	}
	return v, nil
}

func (v *User) UpdateUserNomor(id uint, pa string, db *gorm.DB) (*User, error) {

	if err := db.Debug().Table("users").Where("id = ?", id).Updates(User{
		Nomor_hp: pa,
	}).Error; err != nil {
		return &User{}, err
	}
	return v, nil
}
