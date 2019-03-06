package server

import (
	"fmt"

	validator "gopkg.in/go-playground/validator.v9"
)

// Validator バリデーションのチェックを行う
type Validator struct {
	validator *validator.Validate
}

// NewValidator バリデーターを生成する
func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

// Validate バリデーションをする
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

// ValidationType バリデーションの種類
type ValidationType string

// FieldKey バリデーションする項目名
type FieldKey string

// ValidationResult バリデーション結果
type ValidationResult map[FieldKey]ValidationFiledResult

// ValidationFiledResult フィールドのバリデーションした結果
type ValidationFiledResult []string

// FieldLabels バリデーションした結果の名前
var FieldLabels = map[FieldKey]string{
	"Login":    "ログイン名",
	"Password": "パスワード",
}

// ValidationLabels バリデーションのラベル
var ValidationLabels = map[ValidationType]string{
	"required": "%s を入力してください",
}

// ReportValidation バリデーション結果を分解します
func ReportValidation(errs error) ValidationResult {
	result := ValidationResult{}
	for _, err := range errs.(validator.ValidationErrors) {
		key := FieldKey(err.Field())
		typ := ValidationType(err.Tag())

		// labelize
		fieldLabel, ok := FieldLabels[key]
		if !ok {
			fieldLabel = "undefined"
		}

		validationLabel, ok := ValidationLabels[typ]
		if !ok {
			validationLabel = "undefiend: label: %s"
		}

		// cleate validator list
		label := fmt.Sprintf(validationLabel, fieldLabel)
		if v, ok := result[key]; ok {
			result[key] = append(v, label)
		} else {
			result[key] = []string{label}
		}
	}
	return result
}
