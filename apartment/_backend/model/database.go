package model

import (
	"golang.org/x/crypto/bcrypt"
)

// Admin 管理者
type Admin struct {
	Login    string `gorm:"primary_key"`
	Password string
}

// SetPassword パスワードとして password をハッシュ化する
func (m *Admin) SetPassword(password string) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	m.Password = string(hashed)
}

// ComparePassword ハッシュ化されたパスワードと password を比較する
func (m *Admin) ComperPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(password))
	return err != nil
}

type Apartment struct {
	ID   int `gorm:"primary_key"`
	Name string
}

func (m Apartment) ToViewModel() ApartmentView {
	return ApartmentView{
		ID:   m.ID,
		Name: m.Name,
	}
}
