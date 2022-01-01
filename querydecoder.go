// Populate struct from optional query parameters using 'query' tag.
// It can be used to populate value into primitive variable.
package querydecoder

import (
	"errors"
	"log"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type QueryDecoder interface {
	Decode(target interface{}) error
	DecodeField(key string, target interface{}) error
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

	rv := reflect.ValueOf(target)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("target should be pointer")
	}

	rvElems := rv.Elem()

	rTyp := rvElems.Type()
	noOfFields := rvElems.NumField()

	for i := 0; i < noOfFields; i++ {

		key := rTyp.Field(i).Tag.Get("query")

		if key == "" {
			continue
		}

		if _, ok := q.values[key]; !ok {
			continue
		}

		if err := parseAndSetValue(q.values.Get(key), rvElems.Field(i)); err != nil {
			return err
		}

	}

	return nil

}

func (q *queryDecoder) DecodeField(key string, target interface{}) error {

	rv := reflect.ValueOf(target)

	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("target should be pointer")
	}

	rvElems := rv.Elem()

	if _, ok := q.values[key]; !ok {
		return nil
	}

	if err := parseAndSetValue(q.values.Get(key), rvElems); err != nil {
		return err
	}

	return nil
}

func parseAndSetValue(s string, rv reflect.Value) error {
	switch rv.Kind() {
	case reflect.String:
		rv.SetString(s)
	case reflect.Bool:
		rv.SetBool(strings.ToLower(s) == "true")
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		rv.SetInt(n)
	case reflect.Float32, reflect.Float64:
		n, err := strconv.ParseFloat(s, rv.Type().Bits())
		if err != nil {
			return err
		}
		rv.SetFloat(n)

	case reflect.Ptr:
		if rv.IsNil() {
			rv.Set(reflect.New(rv.Type().Elem()))
		}
		if err := parseAndSetValue(s, rv.Elem()); err != nil {
			return err
		}

	default:
		log.Println(rv.Kind().String(), "is not supported.")
		return errors.New("unsupported type")
	}

	return nil
}
