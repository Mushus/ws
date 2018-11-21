package types

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/Mushus/app/api/value/validation"
)

// String 文字列を表します
type String struct {
	Validation validation.Option
}

// Key は識別子を取得します
func (s String) Key() string {
	return "string"
}

// Cast は値をキャストします
func (s String) Cast(value interface{}) (interface{}, error) {
	v, ok := value.(string)
	if !ok {
		return TypedValue{}, fmt.Errorf("invalid type: %T", value)
	}
	return v, nil
}

// Validate はバリデーションを行います
func (s String) Validate(value interface{}, vs ...validation.Option) validation.Result {
	str, ok := value.(string)
	if !ok {
		return validation.InvalidType
	}

	validations := append([]validation.Option{s.Validation}, vs...)

	for _, v := range validations {
		// NOTE: ここで言う文字数とは「ユニコード文字1個=1文字」。図形素結合子等の結合文字を一文字としてカウントします。
		// 例えばゼロ幅接合子(ZWJ)を含む👨‍👩‍👦‍👦は「父/ZWJ/母/ZWJ/子供/ZWJ/子供」の合字なので7文字扱いになります。
		runeCount := utf8.RuneCountInString(str)

		// 文字列の長さ
		if v.MaxLength > 0 && runeCount > v.MaxLength {
			return validation.InvalidMaxStrLength
		}

		if runeCount < v.MinLength {
			return validation.InvalidMinStrLength
		}

		byteCount := len(str)

		// UTF-8基準でのbyte数
		if v.MaxByte > 0 && byteCount > v.MaxByte {
			// TODO: 名前
			return validation.InvalidMaxStrLength
		}

		if byteCount < v.MinByte {
			// TODO: 名前
			return validation.InvalidMinStrLength
		}

		// フォーマットが正しいか
		if v.Format != nil && !v.Format.MatchString(str) {
			return validation.InvalidFormat
		}

		// NGワードが含まれていないか
		if v.NGWord != nil && isContainWord(str, v.NGWord...) {
			return validation.InvalidNGWord
		}
	}
	return validation.Valid
}

// 部分文字が含まれているかチェックする
func isContainWord(str string, ngWord ...string) bool {
	for _, ng := range ngWord {
		if strings.Contains(str, ng) {
			return true
		}
	}
	return false
}
