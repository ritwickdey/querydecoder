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

		val, err := parseValue(q.values.Get(queryKeyName), fieldValue.Kind())
		if err != nil {
			return err
		}

		fieldValue.Set(reflect.ValueOf(val).Convert(fieldValue.Type()))
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

	value, err := parseValue(q.values.Get(key), rTargetElem.Kind())

	if err != nil {
		return err
	}

	rTargetElem.Set(reflect.ValueOf(value).Convert(rTargetElem.Type()))

	return nil
}

func parseValue(val string, kind reflect.Kind) (interface{}, error) {
	switch kind {
	case reflect.String:
		return val, nil
	case reflect.Bool:
		return strings.ToLower(val) == "true", nil
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := strconv.ParseInt(val, 10, 64)
		return n, err

	case reflect.Float32, reflect.Float64:
		n, err := strconv.ParseFloat(val, 64)
		return n, err
	default:
		return nil, errors.New("unknown type")
	}

}
