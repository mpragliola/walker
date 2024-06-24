package walker_test

import (
	"testing"

	"github.com/mpragliola/walker"
)

func TestBuilder(t *testing.T) {
	w := walker.NewBuilder("myRoot").WithFS(createTestFs()).Build()

	t.Run("type of w", func(t *testing.T) {
		assertStringType(t, w, "*walker.walker")
	})

	t.Run("simple walk", func(t *testing.T) {
		ch, err := w.Walk()

		assertNoError(t, err)
		assertLen(t, walker.Sink(ch), 7)
	})
}
