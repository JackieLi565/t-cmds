package main

import (
	"testing"
)

func TestParseArgs(t *testing.T) {
	testCases := []string{
    "abc def",
    "abc def ghi",
    "abc 'def ghi'",
    "abc \"def ghi\"",
    "abc    def",
    "abc def   ghi",
    "abc  'def   ghi'  jkl",
    "abc  \"def   ghi\"  jkl",
    "  abc  def ",
    "   abc 'def ghi'   ",
    "  abc \"def ghi\"  ",
    "",
	}

	expected := [][]string{
			{"abc", "def"},
			{"abc", "def", "ghi"},
			{"abc", "def ghi"},
			{"abc", "def ghi"},
			{"abc", "def"},
			{"abc", "def", "ghi"},
			{"abc", "def   ghi", "jkl"},
			{"abc", "def   ghi", "jkl"},
			{"abc", "def"},
			{"abc", "def ghi"},
			{"abc", "def ghi"},
	}
	for i, testCase := range testCases {
		result := parseArgs(testCase)

		if len(expected[i]) != len(result) {
			t.Error(result)
			t.Errorf("expected length: [%d] not equal to result length: [%d]", len(expected), len(result))
			return		
		}

		for j := 0; i < len(expected); i++ {
			if expected[i][j] != result[j] {
				t.Errorf("expected: [%s] does not equal result: [%s] at index %d", expected[i], result[i], i)
			}
			return
		}
	}
}