package admin_test

import (
	"testing"

	"github.com/Mushus/apartment/backend/admin"
)

func TestAdminModelCollect(t *testing.T) {
	user1 := &admin.Admin{
		Login: "admin",
	}
	user1.SetPassword("admin")
	if !user1.ComparePassword("admin") {
		t.Fatal("invalid compare password")
	}
}

func TestAdminModelWrong(t *testing.T) {
	user1 := &admin.Admin{
		Login: "admin",
	}
	user1.SetPassword("admin")
	if user1.ComparePassword("hoge") {
		t.Fatal("invalid compare password")
	}
}
