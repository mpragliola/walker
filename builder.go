package walker

import (
	"io/fs"
	"os"
)

type walkBuilder struct {
	fs      fs.FS
	root    string
	filters []WalkFilter
}

func NewBuilder(root string) *walkBuilder {
	return &walkBuilder{
		root: root,
		fs:   os.DirFS(root),
	}
}

func (b *walkBuilder) WithFS(fs fs.FS) *walkBuilder {
	b.fs = fs
	return b
}

func (b *walkBuilder) AddFilter(filter WalkFilter) *walkBuilder {
	b.filters = append(b.filters, filter)
	return b
}

func (b *walkBuilder) AddFilters(filters ...WalkFilter) *walkBuilder {
	b.filters = append(b.filters, filters...)
	return b
}

func (b *walkBuilder) Build() *walker {
	return &walker{
		fs:      b.fs,
		root:    b.root,
		filters: b.filters,
	}
}
