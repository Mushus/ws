package schema

type typ string

const (
	// TypeID ID型
	TypeID typ = TypeVarchar
	// TypePropOrder 並びの型
	TypePropOrder typ = TypeBigint
	// TypeVarchar SQL での VARCHAR 型
	TypeVarchar typ = "VARCHAR"
	// TypeMediumtext SQL での MEDIUMTEXT 型
	TypeMediumtext typ = "MEDIUMTEXT"
	// TypeLongtext SQL での LONGTEXT 型
	TypeLongtext typ = "LONGTEXT"
	// TypeBigint SQL での BIGINT 型
	TypeBigint typ = "BIGINT"
	// TypeDefault デフォルトの型
	TypeDefault typ = "VARCHAR"
)

type size int

const (
	// SizeUnset サイズ未指定
	SizeUnset size = 0
	// SizeID IDのサイズ
	SizeID size = 24
	// SizeVarcharMax は VARCHAR 型の最大サイズです
	SizeVarcharMax size = 65535
	// SizeMediumtextMax は MEDIUMTEXT 型の最大サイズです
	SizeMediumtextMax size = 16777215
)
