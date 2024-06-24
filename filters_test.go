package walker_test

import (
	"testing"

	"github.com/mpragliola/walker"
)

var testData = []string{
	"myFile.txt",
	"abcd",
	"myDir",
	"myDir.md",
	"myDir/aFile.txt",
}

func TestFilters(t *testing.T) {
	for _, tt := range []struct {
		name     string
		filter   walker.WalkFilter
		expected []bool
	}{
		{
			name:     "FilterIdentity()",
			filter:   walker.FilterIdentity(),
			expected: []bool{true, true, true, true, true},
		},
		{
			name:     "Not FilterIdentity() (negation)",
			filter:   walker.Not(walker.FilterIdentity()),
			expected: []bool{false, false, false, false, false},
		},
		{
			name:     "FilterExtensions()",
			filter:   walker.FilterExtensions(".txt", ".md"),
			expected: []bool{true, false, false, true, true},
		},
		{
			name:     "Not FilterExtensions()",
			filter:   walker.Not(walker.FilterExtensions(".txt", ".md")),
			expected: []bool{false, true, true, false, false},
		},
		{
			name:     "FilterRegex()",
			filter:   walker.FilterRegex(".*i.*"),
			expected: []bool{true, false, true, true, true},
		},
		{
			name:     "Not FilterRegex()",
			filter:   walker.Not(walker.FilterRegex(".*i.*")),
			expected: []bool{false, true, false, false, false},
		},
		{
			name:     "FilterStartsWith() without Base()",
			filter:   walker.FilterStartsWith("a"),
			expected: []bool{false, true, false, false, false},
		},
		{
			name:     "FilterStartsWith() with Base()",
			filter:   walker.Base(walker.FilterStartsWith("a")),
			expected: []bool{false, true, false, false, true},
		},
		{
			name:     "Not(Base(FilterRegex()))",
			filter:   walker.Not(walker.Base(walker.FilterRegex("^a.*"))),
			expected: []bool{true, false, true, true, false},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			for i, source := range testData {
				ok, err := tt.filter(source)

				assertNoError(t, err)
				assertEqual(t, ok, tt.expected[i])
			}
		})
	}

}
