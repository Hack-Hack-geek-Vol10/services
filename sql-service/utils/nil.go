package utils

import "reflect"

const (
	NilString = ""
)

func IsInterfaceNil(value interface{}) bool {
	return value == nil || (reflect.ValueOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil())
}

// 空白要素排除
func RmNilString(strings []string) []string {
	result := []string{}
	for _, v := range strings {
		if v != NilString {
			result = append(result, v)
		}
	}
	return result
}
