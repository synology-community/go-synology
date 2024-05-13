package form

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type File struct {
	Name    string `form:"name" url:"name"`
	Content string `form:"content" url:"content"`
}

func NewEncoder() error {
	return nil
}

type FormValues struct {
	*url.Values
}

func (v *FormValues) Encode() ([]byte, error) {
	return Marshal(v)
}

func Marshal(r any) ([]byte, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	v := reflect.Indirect(reflect.ValueOf(r))
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected type struct, got %T", reflect.TypeOf(r).Name())
	}
	n := v.NumField()
	vT := v.Type()
	for i := 0; i < n; i++ {
		formFieldName := strings.ToLower(vT.Field(i).Name)
		formKindName := "field"
		formTags, kindTags := []string{}, []string{}
		if tags, ok := vT.Field(i).Tag.Lookup("form"); ok {
			formTags = strings.Split(tags, ",")
		}
		if tags, ok := vT.Field(i).Tag.Lookup("kind"); ok {
			kindTags = strings.Split(tags, ",")
		}
		if !(vT.Field(i).IsExported() || vT.Field(i).Anonymous || len(formTags) > 0) {
			continue
		}
		if len(formTags) > 0 {
			formFieldName = formTags[0]
		}
		if len(kindTags) > 0 {
			formKindName = kindTags[0]
		}

		// get field type
		switch vT.Field(i).Type.Kind() {
		case reflect.String:
			w.WriteField(formFieldName, v.Field(i).String())
		case reflect.Int:
			w.WriteField(formFieldName, strconv.Itoa(int(v.Field(i).Int())))
		case reflect.Bool:
			w.WriteField(formFieldName, strconv.FormatBool(v.Field(i).Bool()))
		case reflect.Slice:
			slice := v.Field(i)
			switch vT.Field(i).Type.Elem().Kind() {
			case reflect.String:
				res := []string{}
				for iSlice := 0; iSlice < slice.Len(); iSlice++ {
					item := slice.Index(iSlice)
					res = append(res, item.String())
				}
				w.WriteField(formFieldName, "[\""+strings.Join(res, "\",\"")+"\"]")
			case reflect.Int:
				res := []string{}
				for iSlice := 0; iSlice < slice.Len(); iSlice++ {
					item := slice.Index(iSlice)
					res = append(res, strconv.Itoa(int(item.Int())))
				}
				w.WriteField(formFieldName, "["+strings.Join(res, ",")+"]")
			}
		case reflect.Struct:
			// if !vT.Field(i).Anonymous {
			// 	// support only embedded anonymous structs
			// 	continue
			// }

			var fileName string
			var fileReader io.Reader

			embStruct := v.Field(i)
			embStructT := v.Field(i).Type()
			for j := 0; j < embStruct.NumField(); j++ {
				formTags := strings.Split(embStructT.Field(j).Tag.Get("form"), ",")
				fieldName := formTags[0]
				switch embStruct.Field(j).Kind() {
				case reflect.String:
					switch formKindName {
					case "file":
						switch fieldName {
						case "content":
							fileReader = strings.NewReader(embStruct.Field(j).String())
						case "name":
							fileName = embStruct.Field(j).String()
						default:
							w.WriteField(fieldName, embStruct.Field(j).String())
						}
					default:
						w.WriteField(fieldName, embStruct.Field(j).String())
					}
				case reflect.Int:
					w.WriteField(fieldName, strconv.Itoa(int(embStruct.Field(j).Int())))
				}
			}

			if formKindName == "file" {
				if fw, err := w.CreateFormFile(formFieldName, fileName); err != nil {
					return nil, err
				} else {
					if _, err := io.Copy(fw, fileReader); err != nil {
						return nil, err
					}
				}
			}
		}
	}

	w.Close()

	return b.Bytes(), nil
}
