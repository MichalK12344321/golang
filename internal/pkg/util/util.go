package util

import (
	"encoding/json"
	"reflect"
)

func GetSimpleName(obj any) string {
	objVal := reflect.ValueOf(obj)
	for objVal.Kind() == reflect.Ptr || objVal.Kind() == reflect.Interface {
		objVal = objVal.Elem()
	}
	if objVal.IsValid() {
		return objVal.Type().Name()
	}
	return ""
}

func AnyErrors(err ...error) bool {
	for _, v := range err {
		if v != nil {
			return true
		}
	}
	return false
}

func ToJson(source any) ([]byte, error) {
	return json.MarshalIndent(source, "", "  ")
}

func FromJson(target any, data []byte) error {
	return json.Unmarshal(data, target)
}
