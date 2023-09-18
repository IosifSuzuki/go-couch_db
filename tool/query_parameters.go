package tool

import (
	"encoding/json"
	"github.com/IosifSuzuki/go-couch_db/errs"
	"net/url"
	"reflect"
	"strings"
)

const (
	queryKey  = "query"
	omitempty = "omitempty"
)

func BuildUrlValues(value any) (url.Values, error) {
	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)

	urlValues := make(url.Values)

	for i := 0; i < t.NumField(); i++ {
		fieldType := t.Field(i)
		fieldValue := v.Field(i)
		options := fieldType.Tag.Get(queryKey)
		optionsArr := strings.Split(options, ",")
		if len(optionsArr) == 0 {
			return nil, errs.QueryBadConfiguredTagError
		}
		switch fieldValue.Kind() {
		case reflect.String:
			value := fieldValue.String()
			if strings.HasSuffix(options, omitempty) && len(value) == 0 {
				break
			}
			urlValues.Add(optionsArr[0], fieldValue.String())
		case reflect.Int:
			value := fieldValue.Int()
			if strings.HasSuffix(options, omitempty) && value == 0 {
				break
			}
			fieldValueJSON, err := json.Marshal(fieldValue.Int())
			if err != nil {
				return nil, err
			}
			urlValues.Add(optionsArr[0], string(fieldValueJSON))
		case reflect.Bool:
			value := fieldValue.Bool()
			if strings.HasSuffix(options, omitempty) && !value {
				break
			}
			fieldValueJSON, err := json.Marshal(fieldValue.Bool())
			if err != nil {
				return nil, err
			}
			urlValues.Add(optionsArr[0], string(fieldValueJSON))
		}
	}
	return urlValues, nil
}
