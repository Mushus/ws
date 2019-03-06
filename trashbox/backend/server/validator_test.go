package server_test

import (
	"reflect"
	"testing"

	"github.com/Mushus/trashbox/backend/server"
)

func TestValidator(t *testing.T) {
	v := server.NewValidator()
	prm := server.LoginParam{}
	err := v.Validate(prm)
	if err == nil {
		t.Fatal("expect error")
	}

	result := server.ReportValidation(err)
	want := server.ValidationResult{
		"Login":    {"ログイン名 を入力してください"},
		"Password": {"パスワード を入力してください"},
	}
	if !reflect.DeepEqual(result, want) {
		t.Fatalf("result is %#v, expect %#v", result, want)
	}
}
