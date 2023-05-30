package main

import (
	"testing"
)

func TestGetDate(t *testing.T) {
	tests := []struct {
		inputDate string
		want      string
	}{
		{inputDate: "20230130", want: "20230330"},
		{inputDate: "20231231", want: "20240131"},
		{inputDate: "20230228", want: "20230328"},
		{inputDate: "20231031", want: "20231231"},
		{inputDate: "20230331", want: "20230531"},
		{inputDate: "20230531", want: "20230731"},
	}

	for _, test := range tests {

		testName := test.inputDate
		// testName := fmt.Sprint(test.inputDate, " expected result: ", test.want)
		t.Run(testName, func(t *testing.T) {
			result, _ := GetDate(test.inputDate)
			if result != test.want {
				t.Errorf("Expected result: %s, but got: %s", test.want, result)
			}

		})
	}
}

func BenchmarkGetDate(b *testing.B) {
	inputDate := "20230128"

	for n := 0; n < b.N; n++ {
		_, _ = GetDate(inputDate)
	}
}
