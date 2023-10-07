package helper

import (
	"fmt"
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ToBsonD convert regular struct to bson.D
func ToBsonD(src interface{}, dest *bson.D, update bool, parent ...string) {
	var tag = "bson"
	srcValue := reflect.ValueOf(src)
	if srcValue.Kind() != reflect.Struct {
		panic("Required struct")
	}
	total := srcValue.NumField()

	for i := 0; i < total; i++ {
		field := srcValue.Field(i)
		fieldTag := strings.Split(srcValue.Type().Field(i).Tag.Get(tag), ",")[0]

		if !field.IsZero() {
			if field.Kind() == reflect.Struct {
				// struct field category
				var out bson.D
				ToBsonD(field.Interface(), &out, update, fieldTag)
				if !update {
					*dest = append(*dest, primitive.E{Key: fieldTag, Value: out})
					continue
				}
				*dest = append(*dest, out...)
				continue
			}

			var key string = fieldTag
			if len(parent) > 0 && update {
				key = fmt.Sprintf("%s.%s", parent[0], fieldTag)
			}
			*dest = append(*dest, primitive.E{Key: key, Value: field.Interface()})
		}
	}

}
