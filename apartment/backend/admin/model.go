package admin

import (
	"golang.org/x/crypto/bcrypt"
)

// Admin 管理者
type Admin struct {
	Login    string `db:"login"`
	Password string `db:"password"`
}

// SetPassword パスワードとして password をハッシュ化する
func (m *Admin) SetPassword(password string) {
	if m == nil {
		return
	}
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	m.Password = string(hashed)
}

// ComparePassword ハッシュ化されたパスワードと password を比較する
func (m *Admin) ComparePassword(password string) bool {
	// タイミング攻撃強くするために有効かもしれない
	myPassword := ""
	if m != nil {
		myPassword = m.Password
	}
	err := bcrypt.CompareHashAndPassword([]byte(myPassword), []byte(password))
	return err == nil
}
