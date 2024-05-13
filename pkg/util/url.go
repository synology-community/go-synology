package util

import (
	"fmt"
	"mime/multipart"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type Collection interface {
	Add(key string, value string)
}

func MarshalForm(r any, w *multipart.Writer) error {
	add := func(key, value string) {
		w.WriteField(key, value)
	}
	if err := MarshalCollection(r, add); err != nil {
		return err
	}

	return nil
}

func MarshalURL(r any) (url.Values, error) {
	ret := url.Values{}

	add := func(key, value string) {
		ret.Add(key, value)
	}

	if err := MarshalCollection(r, add); err != nil {
		return nil, err
	}

	return ret, nil
}

func MarshalCollection(r any, add func(string, string)) error {
	v := reflect.Indirect(reflect.ValueOf(r))
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("expected type struct, got %T", reflect.TypeOf(r).Name())
	}
	v.FieldByName("")
	n := v.NumField()
	vT := v.Type()
	for i := 0; i < n; i++ {
		urlFieldName := strings.ToLower(vT.Field(i).Name)
		synologyTags := []string{}
		if tags, ok := vT.Field(i).Tag.Lookup("form"); ok {
			synologyTags = strings.Split(tags, ",")
		}
		if !(vT.Field(i).IsExported() || vT.Field(i).Anonymous || len(synologyTags) > 0) {
			continue
		}
		if len(synologyTags) > 0 {
			urlFieldName = synologyTags[0]
		}

		// get field type
		switch vT.Field(i).Type.Kind() {
		case reflect.String:
			add(urlFieldName, v.Field(i).String())
		case reflect.Int:
			add(urlFieldName, strconv.Itoa(int(v.Field(i).Int())))
		case reflect.Bool:
			add(urlFieldName, strconv.FormatBool(v.Field(i).Bool()))
		case reflect.Slice:
			slice := v.Field(i)
			switch vT.Field(i).Type.Elem().Kind() {
			case reflect.String:
				res := []string{}
				for iSlice := 0; iSlice < slice.Len(); iSlice++ {
					item := slice.Index(iSlice)
					res = append(res, item.String())
				}
				add(urlFieldName, "[\""+strings.Join(res, "\",\"")+"\"]")
			case reflect.Int:
				res := []string{}
				for iSlice := 0; iSlice < slice.Len(); iSlice++ {
					item := slice.Index(iSlice)
					res = append(res, strconv.Itoa(int(item.Int())))
				}
				add(urlFieldName, "["+strings.Join(res, ",")+"]")
			}
		case reflect.Struct:
			if !vT.Field(i).Anonymous {
				// support only embedded anonymous structs
				continue
			}
			embStruct := v.Field(i)
			embStructT := v.Field(i).Type()
			for j := 0; j < embStruct.NumField(); j++ {
				synologyTags := strings.Split(embStructT.Field(j).Tag.Get("form"), ",")
				fieldName := synologyTags[0]
				switch embStruct.Field(j).Kind() {
				case reflect.String:
					add(fieldName, embStruct.Field(j).String())
				case reflect.Int:
					add(fieldName, strconv.Itoa(int(embStruct.Field(j).Int())))
				}
			}
		}
	}

	return nil
}
