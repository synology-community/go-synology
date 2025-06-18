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
	Name    string `form:"name"    url:"name"`
	Content string `form:"content" url:"content"`
}

func (f *File) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	*f = File{
		Content: string(data),
	}

	return nil
}

func NewEncoder() error {
	return nil
}

type FormValues struct {
	*url.Values
}

func (v *FormValues) Encode(buf *bytes.Buffer) (*multipart.Writer, int64, error) {
	return Marshal(buf, v)
}

func Marshal(b *bytes.Buffer, input ...any) (*multipart.Writer, int64, error) {
	w := multipart.NewWriter(b)
	defer func() { _ = w.Close() }()
	fileSize := int64(0)

	_ = w.SetBoundary("AaB03x")

	for _, r := range input {

		var v reflect.Value

		if reflect.TypeOf(r).Kind() == reflect.Ptr {
			v = reflect.Indirect(reflect.ValueOf(r))
		} else {
			v = reflect.ValueOf(r)
		}
		if v.Kind() != reflect.Struct {
			return nil, fileSize, fmt.Errorf(
				"expected type struct, got %T",
				reflect.TypeOf(r).Name(),
			)
		}
		n := v.NumField()
		vT := v.Type()
		for i := 0; i < n; i++ {
			field := vT.Field(i)
			fieldType := field.Type

			// if fieldType.Kind() == reflect.Ptr {
			// 	if v.Field(i).IsNil() {
			// 		continue
			// 	}
			// 	fieldType = fieldType.Elem()
			// }

			formFieldName := strings.ToLower(field.Name)
			formKindName := "field"
			formTags, kindTags := []string{}, []string{}
			if tags, ok := field.Tag.Lookup("form"); ok {
				formTags = strings.Split(tags, ",")
			}
			if tags, ok := field.Tag.Lookup("kind"); ok {
				kindTags = strings.Split(tags, ",")
			}
			if !field.IsExported() && !field.Anonymous && len(formTags) <= 0 {
				continue
			}
			if len(formTags) > 0 {
				formFieldName = formTags[0]
				if formFieldName == "-" {
					continue
				}
			}
			if len(kindTags) > 0 {
				formKindName = kindTags[0]
			}

			formKind := fieldType.Kind()

			if field.Name == "File" {
				formKind = reflect.Struct
			}

			// get field type
			switch formKind {
			case reflect.String:
				if err := w.WriteField(formFieldName, v.Field(i).String()); err != nil {
					return nil, fileSize, err
				}
			case reflect.Int:
				if err := w.WriteField(formFieldName, strconv.Itoa(int(v.Field(i).Int()))); err != nil {
					return nil, fileSize, err
				}
			case reflect.Bool:
				if err := w.WriteField(formFieldName, strconv.FormatBool(v.Field(i).Bool())); err != nil {
					return nil, fileSize, err
				}
			case reflect.Slice:
				slice := v.Field(i)
				switch fieldType.Elem().Kind() {
				case reflect.String:
					res := []string{}
					for iSlice := 0; iSlice < slice.Len(); iSlice++ {
						item := slice.Index(iSlice)
						res = append(res, item.String())
					}
					if err := w.WriteField(formFieldName, "[\""+strings.Join(res, "\",\"")+"\"]"); err != nil {
						return nil, fileSize, err
					}
				case reflect.Int:
					res := []string{}
					for iSlice := 0; iSlice < slice.Len(); iSlice++ {
						item := slice.Index(iSlice)
						res = append(res, strconv.Itoa(int(item.Int())))
					}
					if err := w.WriteField(formFieldName, "["+strings.Join(res, ",")+"]"); err != nil {
						return nil, fileSize, err
					}
				}
			case reflect.Struct:
				// if !field.Anonymous {
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
								if err := w.WriteField(fieldName, embStruct.Field(j).String()); err != nil {
									return nil, fileSize, err
								}
							}
						default:
							if err := w.WriteField(fieldName, embStruct.Field(j).String()); err != nil {
								return nil, fileSize, err
							}
						}
					case reflect.Int:
						if err := w.WriteField(fieldName, strconv.Itoa(int(embStruct.Field(j).Int()))); err != nil {
							return nil, fileSize, err
						}
					}
				}

				if formKindName == "file" {
					if fw, err := w.CreateFormFile(formFieldName, fileName); err != nil {
						return nil, fileSize, err
					} else {
						if size, err := io.Copy(fw, fileReader); err != nil {
							return nil, fileSize, err
						} else {
							fileSize = size
						}
					}
				}
			}
		}
	}

	return w, fileSize, nil
}
