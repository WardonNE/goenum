package goenum

import (
	"errors"
	"fmt"
)

var ErrInvalidEnumValue = errors.New("invalid value for enum")

type Enum[T fmt.Stringer] []T

// Values() get all enums
func (this Enum[T]) Values() []T {
	return this
}

// Index() get the index of given enum value
// return -1 if not found
func (this Enum[T]) Index(value T) int {
	for index, item := range this {
		if item.String() == value.String() {
			return index
		}
	}
	return -1
}

// Exists() does the enumeration contain the given value
func (this Enum[T]) Exists(value T) bool {
	return this.Index(value) >= 0
}

// FromString() get enum by given string value
func (this Enum[T]) FromString(value string) (T, error) {
	var empty T
	for _, item := range this {
		if item.String() == value {
			return item, nil
		}
	}
	return empty, ErrInvalidEnumValue
}
