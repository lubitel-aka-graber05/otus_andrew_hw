package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	sl := []string{"77777777777", "88888888888"}
	var err ValidationError
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			User{
				ID:     "1",
				Name:   "AndrewOleshko",
				Age:    101,
				Email:  "fff@fff.com",
				Role:   "Director",
				Phones: sl,
				meta:   nil,
			},
			err.Err,
		},

		{
			App{
				Version: "0.0.1a",
			},
			err.Err,
		},

		{
			Response{
				Code: 202,
				Body: "",
			},
			err.Err,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			t.Log(tt.expectedErr)
			_ = tt
		})
	}
}
