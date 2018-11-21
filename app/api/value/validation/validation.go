package validation

import (
	"regexp"

	"github.com/Mushus/app/api/config"
)

// Option はバリデーションの設定です
//
// define validation name
//
// バリデーションの名前を決める時に注意したほうがいいことがいくらかあります。
// 機能の名前をつけるのではなく、ユースケースに合わせて名前をつけることが重要です。
// 例えば形式が適切かチェックするバリデーションでは regexp より format のほうがより適切です。
type Option struct {
	MaxLength int
	MinLength int
	MaxByte   int
	MinByte   int
	Format    *regexp.Regexp
	NGWord    []string
}

// Result バリデーションの結果を表す値です
type Result string

const (
	// Valid エラーなし
	Valid Result = ""
	// InvalidType タイプが間違ってるエラー
	InvalidType Result = "type"
	// InvalidMaxStrLength 文字数が多いエラー
	InvalidMaxStrLength Result = "max_str_length"
	// InvalidMinStrLength 文字数が少ないエラー
	InvalidMinStrLength Result = "min_str_length"
	// InvalidMaxListLength リストが長すぎるエラー
	InvalidMaxListLength Result = "max_list_length"
	// InvalidMinListLength リストが短すぎるエラー
	InvalidMinListLength Result = "min_list_length"
	// InvalidFormat フォーマットにあっていないエラー
	InvalidFormat Result = "format"
	// InvalidNGWord 使用できない文字が含まれてるエラー
	// NOTE: 改行を含むので、エラー文言は「使用できない文字が含まれています」と出す
	InvalidNGWord Result = "ng_word"
)

// CreateOption 設定ファイルからバリデーションを作成します
func CreateOption(cfg config.Validation) Option {
	return Option{}
}
