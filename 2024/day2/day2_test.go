package main

import (
	"slices"
	"strconv"
	"testing"
)

func TestParseReportsFromPath_Example_ShouldParseAll(t *testing.T) {
	testDataPath := "testdata/example1.txt"
	actual, err := parseReportsFromPath(testDataPath)
	if err != nil {
		t.Fatal(err)
	}
	expected := []report{
		{
			levels: []int{7, 6, 4, 2, 1},
		},
		{
			levels: []int{1, 2, 7, 8, 9},
		},
		{
			levels: []int{9, 7, 6, 2, 1},
		},
		{
			levels: []int{1, 3, 2, 4, 5},
		},
		{
			levels: []int{8, 6, 4, 4, 1},
		},
		{
			levels: []int{1, 3, 6, 7, 9},
		},
	}
	if len(actual) != len(expected) {
		t.Fatalf(`expected len: "%d", got: "%d"`, len(expected), len(actual))
	}
	for idx, r := range expected {
		a := actual[idx]
		if !slices.Equal(r.levels, a.levels) {
			t.Fatalf(`expected levels: "%v", got: "%v"`, r.levels, a.levels)
		}
	}
}

func TestIsReportSafe_Example_ShouldMatchExpected(t *testing.T) {
	testCases := []struct {
		testReport report
		isSafe     bool
	}{
		{
			testReport: report{
				levels: []int{7, 6, 4, 2, 1},
			},
			isSafe: true,
		},
		{
			testReport: report{
				levels: []int{1, 2, 7, 8, 9},
			},
			isSafe: false,
		},
		{
			testReport: report{
				levels: []int{9, 7, 6, 2, 1},
			},
			isSafe: false,
		},
		{
			testReport: report{
				levels: []int{1, 3, 2, 4, 5},
			},
			isSafe: false,
		},
		{
			testReport: report{
				levels: []int{8, 6, 4, 4, 1},
			},
			isSafe: false,
		},
		{
			testReport: report{
				levels: []int{1, 3, 6, 7, 9},
			},
			isSafe: true,
		},
	}
	for idx, tt := range testCases {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			safe := isReportSafe(tt.testReport)
			if safe != tt.isSafe {
				t.Fatalf(`expected report: "%v" safe to be: "%t", got: "%t"`,
					tt.testReport, tt.isSafe, safe)
			}
		})
	}
}

func TestIsReportSafeWithOneRemoved_Example_ShouldMatchExpected(t *testing.T) {
	testCases := []struct {
		testReport report
		isSafe     bool
	}{
		{
			testReport: report{
				levels: []int{7, 6, 4, 2, 1},
			},
			isSafe: true,
		},
		{
			testReport: report{
				levels: []int{1, 2, 7, 8, 9},
			},
			isSafe: false,
		},
		{
			testReport: report{
				levels: []int{9, 7, 6, 2, 1},
			},
			isSafe: false,
		},
		{
			testReport: report{
				levels: []int{1, 3, 2, 4, 5},
			},
			isSafe: true,
		},
		{
			testReport: report{
				levels: []int{8, 6, 4, 4, 1},
			},
			isSafe: true,
		},
		{
			testReport: report{
				levels: []int{1, 3, 6, 7, 9},
			},
			isSafe: true,
		},
	}
	for idx, tt := range testCases {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			safe := isReportSafeWithOneRemoved(tt.testReport)
			if safe != tt.isSafe {
				t.Fatalf(`expected report: "%v" safe to be: "%t", got: "%t"`,
					tt.testReport, tt.isSafe, safe)
			}
		})
	}
}
