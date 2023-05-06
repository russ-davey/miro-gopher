package miro

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

type Parameter map[string]string

type tags struct {
	tag       string
	omitempty bool
	required  bool
}

func encodeQueryParams(queryParams []Parameter) string {
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

func parseQueryTags(v interface{}) []Parameter {
	params := make([]Parameter, 0)
	t := reflect.TypeOf(v)
	value := reflect.ValueOf(v)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if tagStr := field.Tag.Get("query"); tagStr != "" {
			tag := parseTag(tagStr)
			val := fmt.Sprintf("%v", value.Field(i))
			if val == "" && tag.omitempty {
				continue
			}
			params = append(params, Parameter{tag.tag: val})
		}
	}
	return params
}

func parseTag(tagStr string) tags {
	t := tags{}

	keys := strings.Split(tagStr, ",")
	for i, key := range keys {
		if i == 0 {
			t.tag = key
		} else if key == "omitempty" {
			t.omitempty = true
		} else if key == "required" {
			t.omitempty = true
		}
	}
	return t
}
