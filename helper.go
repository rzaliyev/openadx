package openadx

import "strings"

type StrSlice []string

func (s StrSlice) Contains(value string) bool {
	for _, v := range s {
		if strings.EqualFold(v, value) {
			return true
		}
	}
	return false
}

func (s StrSlice) ContainsCaseSensitive(value string) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}
	return false
}
