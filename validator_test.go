package validator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateString(t *testing.T) {
	testcases := []struct {
		name string
		str  string
		min  int
		max  int
		err  error
	}{
		{
			"ok",
			"anthony",
			3,
			25,
			nil,
		},
		{
			"not ok",
			"anto",
			5,
			25,
			fmt.Errorf("must contain characters between  %d - %d ", 5, 25),
		},
	}
	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateString(test.str, test.min, test.max)
			require.EqualValues(t, err, test.err)
		})
	}
}
func TestValidateUserName(t *testing.T) {
	testcases := []struct {
		name string
		str  string
		err  error
	}{
		{
			"ok",
			"anthony456_",
			nil,
		},
		{
			"not ok",
			"Antonio?",
			fmt.Errorf("must contain only lowercase letters, digits, or underscore"),
		},
	}
	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateUserName(test.str)
			require.EqualValues(t, err, test.err)
		})
	}
}
func TestValidatePhone(t *testing.T) {
	testcases := []struct {
		name string
		str  string
		err  error
	}{
		{
			"ok",
			"234567890",
			nil,
		},
		{
			"not ok",
			"3456678i",
			fmt.Errorf("must contain only digits"),
		},
	}
	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			err := ValidatePhone(test.str)
			require.EqualValues(t, err, test.err)
		})
	}
}
func TestValidateName(t *testing.T) {
	testcases := []struct {
		name string
		str  string
		err  error
	}{
		{
			"ok",
			"Anthony",
			nil,
		},
		{
			"not ok",
			"Anthony567",
			fmt.Errorf("must contain only letters, spaces"),
		},
	}
	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateName(test.str)
			require.EqualValues(t, err, test.err)
		})
	}
}

func TestValidatePassword(t *testing.T) {
	testcases := []struct {
		name string
		str  string
		err  error
	}{
		{
			"ok",
			"anthony",
			nil,
		},
		{
			"not ok",
			"anto",
			fmt.Errorf("must contain characters between  %d - %d ", 6, 100),
		},
	}
	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			err := ValidatePassword(test.str)
			require.EqualValues(t, err, test.err)
		})
	}
}
func TestValidateEmail(t *testing.T) {
	testcases := []struct {
		name string
		str  string
		err  error
	}{
		{
			"ok",
			"example@gmail.com",
			nil,
		},
		{
			"not ok",
			"example#gmail.com",
			fmt.Errorf("is not a valid email address"),
		},
	}
	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateEmail(test.str)
			require.EqualValues(t, err, test.err)
		})
	}
}
