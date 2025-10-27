package docling

import (
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"reflect"
	"strings"
)

func multipartEncode(w *multipart.Writer, s any) error {
	v := reflect.Indirect(reflect.ValueOf(s))
	if v.Kind() != reflect.Struct {
		return errors.New("value must be a struct")
	}
	t := v.Type()
	for i := range v.NumField() {
		f := v.Field(i)
		sf := t.Field(i)
		tag := sf.Tag.Get("json")
		if tag == "-" {
			continue
		}
		name, opts := parseTag(tag)
		if name == "" {
			name = strings.ToLower(sf.Name)
		}
		omitEmpty := contains(opts, "omitempty")
		omitZero := contains(opts, "omitzero")
		if (omitEmpty && isEmptyValue(f)) || (omitZero && f.IsZero()) {
			continue
		}
		err := writeValue(w, name, f)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeValue(w *multipart.Writer, name string, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Bool:
		ff, err := w.CreateFormField(name)
		if err != nil {
			return err
		}
		_, err = fmt.Fprint(ff, v.Bool())
		if err != nil {
			return err
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ff, err := w.CreateFormField(name)
		if err != nil {
			return err
		}
		_, err = fmt.Fprint(ff, v.Int())
		if err != nil {
			return err
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		ff, err := w.CreateFormField(name)
		if err != nil {
			return err
		}
		_, err = fmt.Fprint(ff, v.Uint())
		if err != nil {
			return err
		}
	case reflect.Float32, reflect.Float64:
		ff, err := w.CreateFormField(name)
		if err != nil {
			return err
		}
		_, err = fmt.Fprint(ff, v.Float())
		if err != nil {
			return err
		}
	case reflect.String:
		ff, err := w.CreateFormField(name)
		if err != nil {
			return err
		}
		_, err = fmt.Fprint(ff, v.String())
		if err != nil {
			return err
		}
	case reflect.Map, reflect.Struct:
		data, err := json.Marshal(v.Interface())
		if err != nil {
			return err
		}
		ff, err := w.CreateFormField(name)
		if err != nil {
			return err
		}
		_, err = ff.Write(data)
		if err != nil {
			return err
		}
	case reflect.Array, reflect.Slice:
		for j := range v.Len() {
			ff, err := w.CreateFormField(name)
			if err != nil {
				return err
			}
			_, err = fmt.Fprint(ff, v.Index(j).Interface())
			if err != nil {
				return err
			}
		}
	case reflect.Interface:
		return writeValue(w, name, reflect.ValueOf(v.Interface()))
	case reflect.Pointer:
		return writeValue(w, name, reflect.Indirect(v))
	default:
		return fmt.Errorf("unsupported type: %v", v.Kind())
	}
	return nil
}

func parseTag(tag string) (string, string) {
	tag, opt, _ := strings.Cut(tag, ",")
	return tag, opt
}

func contains(options, optionName string) bool {
	if len(options) == 0 {
		return false
	}
	s := options
	for s != "" {
		var name string
		name, s, _ = strings.Cut(s, ",")
		if name == optionName {
			return true
		}
	}
	return false
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.Interface, reflect.Pointer:
		return v.IsZero()
	}
	return false
}
