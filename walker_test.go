package walker_test

import (
	"testing"

	"github.com/mpragliola/walker"
)

func TestWalk(t *testing.T) {

	for _, tt := range []struct {
		name      string
		filters   []walker.WalkFilter
		pathCount int
	}{
		{
			name:      "simple walk",
			filters:   []walker.WalkFilter{},
			pathCount: 7,
		},
		{
			name:      "simple filter",
			filters:   []walker.WalkFilter{walker.FilterIdentity()},
			pathCount: 7,
		},
		{
			name:      "filtering out all paths",
			filters:   []walker.WalkFilter{walker.Not(walker.FilterIdentity())},
			pathCount: 0,
		},
		{
			name:      "filtering all paths with extensions",
			filters:   []walker.WalkFilter{walker.FilterExtensions(".txt", ".md")},
			pathCount: 6,
		},
		{
			name:      "filtering all paths with not regexp",
			filters:   []walker.WalkFilter{walker.Not(walker.FilterRegex(".*i.*"))},
			pathCount: 6,
		},
		{
			name:      "filtering all paths with regexp without base",
			filters:   []walker.WalkFilter{walker.FilterRegex("^_.*")},
			pathCount: 0,
		},
		{
			name:      "filtering all paths with regexp on base",
			filters:   []walker.WalkFilter{walker.Base(walker.FilterRegex("^_.*"))},
			pathCount: 2,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			w := walker.NewBuilder("myRoot").
				WithFS(createTestFs()).
				AddFilters(tt.filters...).
				Build()

			ch, err := w.Walk()

			assertNoError(t, err)
			assertLen(t, walker.Sink(ch), tt.pathCount)
		})
	}
}
