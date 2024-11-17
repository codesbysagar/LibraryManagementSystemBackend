package main

import (
	"errors"
	"math"
	"testing"
)

func TestIdGenerator(t *testing.T) {
	testCases := []struct {
		name  string
		input int
		err   error
	}{
		{
			name:  "Zero",
			input: 0,
			err:   errors.New("digits limit is not valid. It should be greater than 0"),
		},
		{
			name:  "Greater than Zero",
			input: 5,
			err:   nil,
		},
		{
			name:  "Smaller than Zero",
			input: -4,
			err:   errors.New("digits limit is not valid. It should be greater than 0"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result, err := IdGenerator(tt.input)
			upperBound := int(math.Pow10(tt.input))
			lowerBound := int(math.Pow10(tt.input - 1))
			if err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("expected error %v got error %v", tt.err, err)
				}
			} else if result < lowerBound || result >= upperBound {
				t.Errorf("Generated Id %d is out of bound of digits : %d (expected between %d -- %d)", result, tt.input, lowerBound, upperBound)
			}

		})
	}
}

func TestMemberValidator(t *testing.T) {
	testCases := []struct {
		name  string
		input MemberStruct
		want  error
	}{
		{
			name: "Valid",
			input: MemberStruct{
				FullName: "Sagar Sharma",
				Contact:  1212121212,
				Email:    "sagar.sharma@example.com",
				Password: "thisispass",
			},
			want: nil,
		},
		{
			name: "Missing fullName",
			input: MemberStruct{
				FullName: "",
				Contact:  1212121212,
				Email:    "sagar.sharma@example.com",
				Password: "thisispass",
			},
			want: errors.New("fullName is missing"),
		},
		{
			name: "Missing Contact",
			input: MemberStruct{
				FullName: "Sagar Sharma",
				Contact:  12121,
				Email:    "sagar.sharma@example.com",
				Password: "thisispass",
			},
			want: errors.New("contact number should be of 10 digits only"),
		},
		{
			name: "Invalid Email",
			input: MemberStruct{
				FullName: "Sagar Sharma",
				Contact:  1212121212,
				Email:    "sagar.sharma@example.com",
				Password: "thisispass",
			},
			want: errors.New("invalid email address"),
		},
		{
			name: "Invalid Password",
			input: MemberStruct{
				FullName: "Sagar Sharma",
				Contact:  1212121212,
				Email:    "sagar.sharma@example.com",
				Password: "ass",
			},
			want: errors.New("password cannot be empty or less than 8 digits"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			err := MemberValidator(tt.input)
			if err != nil {
				if err.Error() != tt.want.Error() {
					t.Errorf("expected error %v got error %v", tt.want, err)
				}
			}
		})
	}
}
