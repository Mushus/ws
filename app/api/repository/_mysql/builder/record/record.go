package record

import "errors"

// Record とあるいちレコードを表します
type Records []SortedValue

// ScanValue 指定したインデックスのカラムを集めます
func (r Records) ScanValue(index int) ([]Value, error) {
	values := make(SortedValue, len(r))
	for i, value := range r {
		if len(value) <= index {
			return nil, errors.New("index out of range")
		}
		values[i] = value[index]
	}
	return values, nil
}
