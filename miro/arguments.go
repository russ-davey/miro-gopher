package miro

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

type Arguments map[string]string

func EncodeQueryParams(queryParams []Arguments) string {
	values := url.Values{}
	for _, params := range queryParams {
		for key, value := range params {
			values.Add(key, value)
		}
	}
	if len(values) == 0 {
		return ""
	}
	return "?" + values.Encode()
}

func ParseQueryTags(v interface{}) []Arguments {
	args := make([]Arguments, 0)
	t := reflect.TypeOf(v)
	value := reflect.ValueOf(v)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if tag := field.Tag.Get("query"); tag != "" {
			val := fmt.Sprintf("%v", value.Field(i))
			if val == "" {
				if strings.Contains(tag, "omitempty") {
					continue
				}
				val = ""
			}

			key := strings.Replace(tag, ",omitempty", "", 1)
			args = append(args, Arguments{key: val})
		}
	}
	return args
}
