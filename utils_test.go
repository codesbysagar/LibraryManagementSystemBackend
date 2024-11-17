package main

import (
	"errors"
	"math"
	"testing"
)

func TestIdGenerator(t *testing.T) {
	testCases := []struct{
		name string
		input int
		err error
	}{
		{
			name: "Zero",
			input: 0,
			err: errors.New("digits limit is not valid. It should be greater than 0"),
		},
		{
			name: "Greater than Zero",
			input: 5,
			err: nil,
		},
		{
			name: "Smaller than Zero",
			input: -4,
			err: errors.New("digits limit is not valid. It should be greater than 0"),
		},
	}

	for _, tt:= range testCases{
		t.Run(tt.name,func(t *testing.T) {
			result, err := IdGenerator(tt.input)
			upperBound := int(math.Pow10(tt.input))
			lowerBound := int(math.Pow10(tt.input-1))
			if err!= nil {
				if err.Error() != tt.err.Error(){
					t.Errorf("expected error %v got error %v", tt.err, err)
				}
			}else if result<lowerBound || result>=upperBound {
				t.Errorf("Generated Id %d is out of bound of digits : %d (expected between %d -- %d)", result, tt.input, lowerBound, upperBound)
			}
			
		})
	}
}
