package main

import (
	"reflect"
	"testing"
)

type testInputOutput struct {
	input         string
	output        Params
	isFirstValid  bool
	isSecondValid bool
}

var samples = map[string]testInputOutput{
	"first":  {input: "1-3 a: abcde", output: Params{firstNum: 1, secondNum: 3, letter: 'a', password: "abcde"}, isFirstValid: true, isSecondValid: true},
	"second": {input: "1-3 b: cdefg", output: Params{firstNum: 1, secondNum: 3, letter: 'b', password: "cdefg"}, isFirstValid: false, isSecondValid: false},
	"third":  {input: "2-9 c: ccccccccc", output: Params{firstNum: 2, secondNum: 9, letter: 'c', password: "ccccccccc"}, isFirstValid: true, isSecondValid: false},
}

func TestGetParams(t *testing.T) {
	for name, sample := range samples {
		t.Run(name, func(t *testing.T) {
			got := GetParams(sample.input)
			want := sample.output
			if reflect.DeepEqual(got, want) == false {
				t.Errorf("\ngiven %s\ngot: {min: %d, max: %d, letter: %c}\nwant: {min: %d, max: %d, letter: %c}", sample.input, got.firstNum, got.secondNum, got.letter, want.firstNum, want.secondNum, want.letter)
			}
		})
	}
}

func TestValidatePassword(t *testing.T) {
	for name, sample := range samples {
		t.Run(name, func(t *testing.T) {
			got := ValidatePassword(sample.output)
			want := sample.isFirstValid
			if got != want {
				t.Errorf("\ngiven %s\ngot: %t\nwant: %t", sample.input, got, want)
			}
		})
	}

}

func TestValidatePasswordTwo(t *testing.T) {
	for name, sample := range samples {
		t.Run(name, func(t *testing.T) {
			got := ValidatePasswordTwo(sample.output)
			want := sample.isSecondValid
			if got != want {
				t.Errorf("\ngiven %s\ngot: %t\nwant: %t", sample.input, got, want)
			}
		})
	}

}
