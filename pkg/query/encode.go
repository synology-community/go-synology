// Package query implements encoding of structs into URL query parameters.
//
// As a simple example:
//
//	type Options struct {
//		Query   string `url:"q"`
//		ShowAll bool   `url:"all"`
//		Page    int    `url:"page"`
//	}
//
//	opt := Options{ "foo", true, 2 }
//	v, _ := query.Values(opt)
//	fmt.Print(v.Encode()) // will output: "q=foo&all=true&page=2"
//
// The exact mapping between Go values and url.Values is described in the
// documentation for the Values() function.
package query

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var timeType = reflect.TypeOf(time.Time{})

var encoderType = reflect.TypeOf(new(Encoder)).Elem()

// Encoder is an interface implemented by any type that wishes to encode
// itself into URL values in a non-standard way.
type Encoder interface {
	EncodeValues(key string, v *url.Values) error
}

// Values returns the url.Values encoding of v.
//
// Values expects to be passed a struct, and traverses it recursively using the
// following encoding rules.
//
// Each exported struct field is encoded as a URL parameter unless
//
//   - the field's tag is "-", or
//   - the field is empty and its tag specifies the "omitempty" option
//
// The empty values are false, 0, any nil pointer or interface value, any array
// slice, map, or string of length zero, and any type (such as time.Time) that
// returns true for IsZero().
//
// The URL parameter name defaults to the struct field name but can be
// specified in the struct field's tag value.  The "url" key in the struct
// field's tag value is the key name, followed by an optional comma and
// options.  For example:
//
//	// Field is ignored by this package.
//	Field int `url:"-"`
//
//	// Field appears as URL parameter "myName".
//	Field int `url:"myName"`
//
//	// Field appears as URL parameter "myName" and the field is omitted if
//	// its value is empty
//	Field int `url:"myName,omitempty"`
//
//	// Field appears as URL parameter "Field" (the default), but the field
//	// is skipped if empty.  Note the leading comma.
//	Field int `url:",omitempty"`
//
//	// Field appears as URL parameter "myName" and the field is marshaled as
//	// a JSON string.
//	Field []int `url:"myName,json"`
//
//	// Field appears as URL parameter "myName" and the field is marshaled as
//	// a JSON string, quoted.
//	Field []string `url:"myName,json,quoted"`
//
// For encoding individual field values, the following type-dependent rules
// apply:
//
// Boolean values default to encoding as the strings "true" or "false".
// Including the "int" option signals that the field should be encoded as the
// strings "1" or "0".
//
// time.Time values default to encoding as RFC3339 timestamps.  Including the
// "unix" option signals that the field should be encoded as a Unix time (see
// time.Unix()).  The "unixmilli" and "unixnano" options will encode the number
// of milliseconds and nanoseconds, respectively, since January 1, 1970 (see
// time.UnixNano()).  Including the "layout" struct tag (separate from the
// "url" tag) will use the value of the "layout" tag as a layout passed to
// time.Format.  For example:
//
//	// Encode a time.Time as YYYY-MM-DD
//	Field time.Time `layout:"2006-01-02"`
//
// Slice and Array values default to encoding as multiple URL values of the
// same name.  Including the "comma" option signals that the field should be
// encoded as a single comma-delimited value.  Including the "space" option
// similarly encodes the value as a single space-delimited string. Including
// the "semicolon" option will encode the value as a semicolon-delimited string.
// Including the "brackets" option signals that the multiple URL values should
// have "[]" appended to the value name. "numbered" will append a number to
// the end of each incidence of the value name, example:
// name0=value0&name1=value1, etc.  Including the "del" struct tag (separate
// from the "url" tag) will use the value of the "del" tag as the delimiter.
// For example:
//
//	// Encode a slice of bools as ints ("1" for true, "0" for false),
//	// separated by exclamation points "!".
//	Field []bool `url:",int" del:"!"`
//
// Anonymous struct fields are usually encoded as if their inner exported
// fields were fields in the outer struct, subject to the standard Go
// visibility rules.  An anonymous struct field with a name given in its URL
// tag is treated as having that name, rather than being anonymous.
//
// Non-nil pointer values are encoded as the value pointed to.
//
// Nested structs have their fields processed recursively and are encoded
// including parent fields in value names for scoping. For example,
//
//	"user[name]=acme&user[addr][postcode]=1234&user[addr][city]=SFO"
//
// All other values are encoded using their default string representation.
//
// Multiple fields that encode to the same URL parameter name will be included
// as multiple URL values of the same name.
func Values(v any) (url.Values, error) {
	values := make(url.Values)
	val := reflect.ValueOf(v)
	for val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return values, nil
		}
		val = val.Elem()
	}

	if v == nil {
		return values, nil
	}

	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("query: Values() expects struct input. Got %v", val.Kind())
	}

	err := reflectValue(values, val, "")
	return values, err
}

// reflectValue populates the values parameter from the struct fields in val.
// Embedded structs are followed recursively (using the rules defined in the
// Values function documentation) breadth-first.
func reflectValue(values url.Values, val reflect.Value, scope string) error {
	var embedded []reflect.Value

	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		sf := typ.Field(i)
		if sf.PkgPath != "" && !sf.Anonymous { // unexported
			continue
		}

		sv := val.Field(i)
		tag := sf.Tag.Get("url")
		if tag == "-" {
			continue
		}
		name, opts := parseTag(tag)

		if name == "" {
			if sf.Anonymous {
				v := reflect.Indirect(sv)
				if v.IsValid() && v.Kind() == reflect.Struct {
					// save embedded struct for later processing
					embedded = append(embedded, v)
					continue
				}
			}

			name = sf.Name
		}

		if scope != "" {
			name = scope + "[" + name + "]"
		}

		if opts.Contains("omitempty") && isEmptyValue(sv) {
			continue
		}

		if sv.Type().Implements(encoderType) {
			// if sv is a nil pointer and the custom encoder is defined on a non-pointer
			// method receiver, set sv to the zero value of the underlying type
			if !reflect.Indirect(sv).IsValid() && sv.Type().Elem().Implements(encoderType) {
				sv = reflect.New(sv.Type().Elem())
			}

			if m, ok := sv.Interface().(Encoder); ok {
				if err := m.EncodeValues(name, &values); err != nil {
					return err
				}
				continue
			} else {
				return fmt.Errorf("query: Values() expects Encoder interface. Got %v", sv.Type())
			}
		}

		// recursively dereference pointers. break on nil pointers
		for sv.Kind() == reflect.Ptr {
			if sv.IsNil() {
				break
			}
			sv = sv.Elem()
		}

		if opts.Contains("json") {
			encoded, err := json.Marshal(sv.Interface())
			if err != nil {
				return err
			}

			s := string(encoded)

			if opts.Contains("quoted") {
				s = strconv.Quote(s)
			}

			values.Add(name, s)
			continue
		}

		if sv.Kind() == reflect.Slice || sv.Kind() == reflect.Array {
			if sv.Len() == 0 {
				// skip if slice or array is empty
				continue
			}

			var del string
			if opts.Contains("comma") {
				del = ","
			} else if opts.Contains("space") {
				del = " "
			} else if opts.Contains("semicolon") {
				del = ";"
			} else if opts.Contains("brackets") {
				name = name + "[]"
			} else {
				del = sf.Tag.Get("del")
			}

			if del != "" {
				s := new(bytes.Buffer)
				first := true
				for i := 0; i < sv.Len(); i++ {
					if first {
						first = false
					} else {
						s.WriteString(del)
					}
					s.WriteString(valueString(sv.Index(i), opts, sf))
				}
				values.Add(name, s.String())
			} else {
				for i := 0; i < sv.Len(); i++ {
					k := name
					if opts.Contains("numbered") {
						k = fmt.Sprintf("%s%d", name, i)
					}
					values.Add(k, valueString(sv.Index(i), opts, sf))
				}
			}
			continue
		}

		if sv.Type() == timeType {
			values.Add(name, valueString(sv, opts, sf))
			continue
		}

		if sv.Kind() == reflect.Struct {
			if err := reflectValue(values, sv, name); err != nil {
				return err
			}
			continue
		}

		values.Add(name, valueString(sv, opts, sf))
	}

	for _, f := range embedded {
		if err := reflectValue(values, f, scope); err != nil {
			return err
		}
	}

	return nil
}

func valueString(v reflect.Value, opts tagOptions, sf reflect.StructField) string {
	if opts.Contains("quoted") {
		return strconv.Quote(valueStringInner(v, opts, sf))
	} else {
		return valueStringInner(v, opts, sf)
	}
}

// valueString returns the string representation of a value.
func valueStringInner(v reflect.Value, opts tagOptions, sf reflect.StructField) string {
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return ""
		}
		v = v.Elem()
	}

	if v.Kind() == reflect.Bool && opts.Contains("int") {
		if v.Bool() {
			return "1"
		} else {
			return "0"
		}
	}

	if v.Type() == timeType {
		if t, ok := v.Interface().(time.Time); ok {
			if opts.Contains("unix") {
				return strconv.FormatInt(t.Unix(), 10)
			}
			if opts.Contains("unixmilli") {
				return strconv.FormatInt((t.UnixNano() / 1e6), 10)
			}
			if opts.Contains("unixnano") {
				return strconv.FormatInt(t.UnixNano(), 10)
			}
			if layout := sf.Tag.Get("layout"); layout != "" {
				return t.Format(layout)
			}
			return t.Format(time.RFC3339)
		} else {
			return ""
		}
	}

	return fmt.Sprint(v.Interface())
}

// isEmptyValue checks if a value should be considered empty for the purposes
// of omitting fields with the "omitempty" option.
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}

	type zeroable interface {
		IsZero() bool
	}

	if z, ok := v.Interface().(zeroable); ok {
		return z.IsZero()
	}

	return false
}

// tagOptions is the string following a comma in a struct field's "url" tag, or
// the empty string. It does not include the leading comma.
type tagOptions []string

// parseTag splits a struct field's url tag into its name and comma-separated
// options.
func parseTag(tag string) (string, tagOptions) {
	s := strings.Split(tag, ",")
	return s[0], s[1:]
}

// Contains checks whether the tagOptions contains the specified option.
func (o tagOptions) Contains(option string) bool {
	for _, s := range o {
		if s == option {
			return true
		}
	}
	return false
}
