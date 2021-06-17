package env2struct

import (
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Options struct {
	Prefix    string
	Separator string
}

func Parse(i interface{}, options ...Options) error {
	// get options
	opt := Options{Separator: "_"}
	if len(options) > 0 {
		if options[0].Prefix != "" {
			opt.Prefix = options[0].Prefix
		}
		if options[0].Separator != "" {
			opt.Separator = options[0].Separator
		}
	}

	v := reflect.ValueOf(i)
	t := reflect.TypeOf(i)

	return parseField(v, t, -1, opt)
}

func parseField(v reflect.Value, t reflect.Type, i int, options Options) error {
	if v.Kind() != reflect.Ptr {
		return errors.New("not a pointer value")
	}
	f := reflect.StructField{}
	if i != -1 {
		f = t.Field(i)
	}
	v = reflect.Indirect(v)
	fieldEnv, exists := f.Tag.Lookup("env")
	env := os.Getenv(options.Prefix + fieldEnv)
	if exists && env != "" {
		switch v.Kind() {
		case reflect.String:
			v.SetString(env)
		case reflect.Bool:
			boolValue, err := strconv.ParseBool(env)
			if err != nil {
				return err
			}
			v.SetBool(boolValue)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intValue, err := strconv.ParseInt(env, 10, 0)
			if err != nil {
				return err
			}
			v.SetInt(intValue)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			intValue, err := strconv.ParseUint(env, 10, 0)
			if err != nil {
				return err
			}
			v.SetUint(intValue)
		case reflect.Float32, reflect.Float64:
			floatValue, err := strconv.ParseFloat(env, 64)
			if err != nil {
				return err
			}
			v.SetFloat(floatValue)
		}
	}
	if v.Kind() == reflect.Struct {
		for j := 0; j < v.NumField(); j++ {
			opt := Options{
				Prefix:    strings.TrimPrefix(options.Prefix+fieldEnv+options.Separator, "_"),
				Separator: options.Separator,
			}
			err := parseField(v.Field(j).Addr(), v.Type(), j, opt)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
