package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var sb strings.Builder
	for i := 0; i < len(v); i++ {
		if i == len(v)-1 {
			_, _ = fmt.Fprint(&sb, v[i])
			continue
		}
		_, _ = fmt.Fprintln(&sb, v[i])
	}
	return sb.String()
}

func Validate(v interface{}) error {
	var err ValidationError
	var containErr ValidationErrors
	var ErrFieldType = errors.New("failed validation: Invalid field type")

	val := reflect.ValueOf(v)
	tp := val.Type()

	if val.Kind() != reflect.Struct {
		err.Field = tp.Name()
		err.Err = errors.New("unexpected struct")
		containErr = append(containErr, err)
	}

	for i := 0; i < val.NumField(); i++ {
		tag := tp.Field(i).Tag.Get("validate")
		body := val.Field(i)

		if tp.Field(i).Tag == "" {
			continue
		}

		switch tp.Field(i).Name {
		case "ID":
			var l int
			_, _ = fmt.Sscanf(tag, "len:%d", &l)
			if val.Field(i).Kind() != reflect.String {
				err.Field = tp.Field(i).Name
				err.Err = ErrFieldType
				containErr = append(containErr, err)
				continue
			}
			if body.Len() != l {
				err.Field = tp.Field(i).Name
				err.Err = errors.New("failed validation: Invalid length")
				containErr = append(containErr, err)
			}
		case "Age":
			if val.Field(i).Kind() != reflect.Int {
				err.Field = tp.Field(i).Name
				err.Err = ErrFieldType
				containErr = append(containErr, err)
				continue
			}

			var min, max int64
			_, _ = fmt.Sscanf(tag, "min:%d|max:%d", &min, &max)

			if body.Int() < min || body.Int() > max {
				err.Field = tp.Field(i).Name
				err.Err = errors.New("failed validation: Invalid age")
				containErr = append(containErr, err)
			}
		case "Email":
			if val.Field(i).Kind() != reflect.String {
				err.Field = tp.Field(i).Name
				err.Err = ErrFieldType
				containErr = append(containErr, err)
				continue
			}
			var mail = regexp.MustCompile(`(?m)^\w+@\w+\.\w+$`)
			if !mail.MatchString(tag) {
				err.Field = tp.Field(i).Name
				err.Err = errors.New("failed validation: Invalid email format")
				containErr = append(containErr, err)
			}
		case "Phones":
			if val.Field(i).Kind() != reflect.Slice {
				err.Field = tp.Field(i).Name
				err.Err = ErrFieldType
				containErr = append(containErr, err)
				continue
			}
			var l int
			if body.Len() != l {
				err.Field = tp.Field(i).Name
				err.Err = errors.New("failed validation: Invalid length")
				containErr = append(containErr, err)
			}
		case "Version":
			var l int
			if val.Field(i).Kind() != reflect.String {
				err.Field = tp.Field(i).Name
				err.Err = ErrFieldType
				containErr = append(containErr, err)
				continue
			}
			if body.Len() != l {
				err.Field = tp.Field(i).Name
				err.Err = errors.New("failed validation: Invalid length")
				containErr = append(containErr, err)
			}
		case "Code":
			if val.Field(i).Kind() != reflect.Int {
				err.Field = tp.Field(i).Name
				err.Err = ErrFieldType
				containErr = append(containErr, err)
				continue
			}
			var a, b, c int64
			_, _ = fmt.Sscanf(tag, "in:%d,%d,%d", &a, &b, &c)
			if body.Int() != a || body.Int() != b || body.Int() != c {
				err.Field = tp.Field(i).Name
				err.Err = errors.New("failed validation: Invalid response")
				containErr = append(containErr, err)
			}
		}
	}

	return containErr
}
