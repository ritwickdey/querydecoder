package querydecoder

import (
	"errors"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type QueryDecoder interface {
	Decode(target interface{}) error
	DecodeField(key string, defaultValue interface{}, target interface{}) error
}

type queryDecoder struct {
	values url.Values
}

func New(values url.Values) QueryDecoder {
	return &queryDecoder{
		values: values,
	}
}

func (q *queryDecoder) Decode(target interface{}) error {

	rVal := reflect.ValueOf(target)
	if rVal.Kind() != reflect.Ptr || rVal.IsNil() {
		return errors.New("target should be pointer")
	}

	rType := reflect.TypeOf(target)

	elemsVal := rVal.Elem()
	elemsType := rType.Elem()

	noOfFields := elemsType.NumField()

	for i := 0; i < noOfFields; i++ {
		queryKeyName := elemsType.Field(i).Tag.Get("query")
		if queryKeyName == "" {
			continue
		}

		if _, ok := q.values[queryKeyName]; !ok {
			continue
		}

		fieldValue := elemsVal.Field(i)

		err := parseAndSetValue(q.values.Get(queryKeyName), &fieldValue)
		if err != nil {
			return err
		}

	}

	return nil

}

func (q *queryDecoder) DecodeField(key string, defaultValue interface{}, target interface{}) error {

	rTarget := reflect.ValueOf(target)

	if rTarget.Kind() != reflect.Ptr || rTarget.IsNil() {
		return errors.New("target should be pointer")
	}

	rTargetElem := rTarget.Elem()

	if _, ok := q.values[key]; !ok {
		rTargetElem.Set(reflect.ValueOf(defaultValue))
		return nil
	}

	err := parseAndSetValue(q.values.Get(key), &rTargetElem)

	if err != nil {
		return err
	}

	return nil
}

func parseAndSetValue(val string, rVal *reflect.Value) error {
	switch rVal.Kind() {
	case reflect.String:
		rVal.SetString(val)
	case reflect.Bool:
		rVal.SetBool(strings.ToLower(val) == "true")
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return err
		}
		rVal.SetInt(n)
	case reflect.Float32, reflect.Float64:
		n, err := strconv.ParseFloat(val, rVal.Type().Bits())
		if err != nil {
			return err
		}
		rVal.SetFloat(n)
	default:
		return errors.New("unknown type")
	}

	return nil
}
