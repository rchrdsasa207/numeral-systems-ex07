package main

import "testing"

func TestNumToBin(t *testing.T) {
	for _, tc := range []struct {
		name   string
		num    int64
		result string
	}{

		{"1", 0, "0"},
		{"zero1", 1, "1"},
		{"zero2", 8, "1000"},
		{"zero3", -1, "1111111111111111111111111111111111111111111111111111111111111111"},
		{"zero4", -1, "1111111111111111111111111111111111111111111111111111111111111111"},
		{"zero5", 1000_0000_0000_0000, "11100011010111111010100100110001101000000000000000"},
		{"zero6", -1000_0000_0000_0001, "1111111111111100011100101000000101011011001110010111111111111111"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := NumToBin(tc.num); got != tc.result {
				t.Errorf("LongAdd(%v) = %v, want %v", tc.num, got, tc.result)
			}
		})
	}
}

func TestSum(t *testing.T) {
	for _, tc := range []struct {
		a      string
		b      string
		result string
	}{
		{"1", "11", "100"},
		{"11001" /* 25 */, "1111111111111111111111111111111111111111111111111111111111110010" /* -14 */, "1011" /* 11 */},
		{"1111111111111111111111111111111111111111111111111111111111110110", "111" /* 7 */, "1111111111111111111111111111111111111111111111111111111111111101" /* -3 */},
		{"11" /* 3 */, "1111111111111111111111111111111111111111111111111111111111111101" /* -3 */, "0" /* 0 */},
	} {
		t.Run("", func(t *testing.T) {
			if got := BinNumSum(tc.a, tc.b); got != tc.result {
				t.Errorf("LongAdd(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.result)
			}
		})
	}
}