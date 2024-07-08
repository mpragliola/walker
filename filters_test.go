package walker_test

import (
	"testing"

	"github.com/mpragliola/walker"
)

// Just a series of paths ...
var testData = []string{
	"myFile.txt",
	"abcd",
	"myDir",
	"myDir.md",
	"myDir/aFile.txt",
	"myDir/other.md",
}

func TestFilters(t *testing.T) {
	for name, tt := range map[string]struct {
		filter   walker.WalkFilter
		expected []bool
	}{
		"FilterIdentity()": {
			filter:   walker.FilterIdentity(),
			expected: []bool{true, true, true, true, true, true},
		},
		"Not FilterIdentity() (negation)": {
			filter:   walker.Not(walker.FilterIdentity()),
			expected: []bool{false, false, false, false, false, false},
		},
		"Filter extensions .txt and .md": {
			filter:   walker.FilterExtensions(".txt", ".md"),
			expected: []bool{true, false, false, true, true, true},
		},
		"Filter extensions that are not .txt or .md": {
			filter:   walker.Not(walker.FilterExtensions(".txt", ".md")),
			expected: []bool{false, true, true, false, false, false},
		},
		"Filter using a regex": {
			filter:   walker.FilterRegex(".*i.*"),
			expected: []bool{true, false, true, true, true, true},
		},
		"Filter using a regex and Base()": {
			filter:   walker.Base(walker.FilterRegex(".*i.*")),
			expected: []bool{true, false, true, true, true, false},
		},
		"Filter excluding regex": {
			filter:   walker.Not(walker.FilterRegex(".*i.*")),
			expected: []bool{false, true, false, false, false, false},
		},
		"FilterStartsWith() without Base()": {
			filter:   walker.FilterStartsWith("a"),
			expected: []bool{false, true, false, false, false, false},
		},
		"FilterStartsWith() with Base()": {
			filter:   walker.Base(walker.FilterStartsWith("a")),
			expected: []bool{false, true, false, false, true, false},
		},
		"Not(Base(FilterRegex()))": {
			filter:   walker.Not(walker.Base(walker.FilterRegex("^a.*"))),
			expected: []bool{true, false, true, true, false, true},
		},
	} {
		t.Run(name, func(t *testing.T) {
			for i, source := range testData {
				ok, err := tt.filter(source)

				assertNoError(t, err)
				assertEqual(t, ok, tt.expected[i])
			}
		})
	}

}
