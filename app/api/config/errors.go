package config

import (
	"errors"
)

var (
	// ErrEntitiesIsRequired は "entities" 必須エラーを表します
	ErrEntitiesIsRequired = errors.New(`"entities" is required`)
)
