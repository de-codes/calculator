package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

type testCase struct {
	a, b float64
	want float64
}

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 21, b: 12, want: 33},
		{a: 25, b: 2, want: 27},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f,%f):want %f , got %f", tc.a, tc.b, tc.want, got)
		}
	}

}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 24, b: 2, want: 22},
		{a: 2, b: 12, want: -10},
		{a: 25, b: 2, want: 23},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f,%f):want %f , got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 24, b: 2, want: 48},
		{a: 1, b: 12, want: 12},
		{a: 25, b: -2, want: -50},
		{a: 25, b: 0, want: 0},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f,%f):want %f , got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 1},
		{a: 10, b: 2, want: 5},
		{a: -1, b: -1, want: 1},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)

		if err != nil {
			t.Fatalf("Divide(%f,%f): want no error for valid input, got %v", tc.a, tc.b, err)
		}
		if tc.want != got {
			t.Errorf("Multiply(%f,%f):want %f , got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	type testCase struct {
		input float64
		want  float64
	}
	testCases := []testCase{
		{input: 4, want: 2},
		{input: 2, want: 1.41421356237},
		{input: 25, want: 5},
		{input: 1.4, want: 1.2},
	}

	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.input)

		if err != nil {
			t.Fatalf("Sqrt(%f): want no error for valid input, got %v", tc.input, err)
		}
		if !closeEnough(tc.want, got, 0.1) {
			t.Errorf("Sqrt(%f):want %f , got %f", tc.input, tc.want, got)
		}
	}
}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()

	_, err := calculator.Divide(1, 0)

	if err == nil {
		t.Errorf("Divide(1,0): want error for invalid input, got %v", err)
	}
}

func TestSqrtInvalid(t *testing.T) {
	t.Parallel()

	_, err := calculator.Sqrt(-1)

	if err == nil {
		t.Errorf("Sqrt(-1): want error for invalid input, got %v", err)
	}
}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}
