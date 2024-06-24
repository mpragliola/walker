# walker

## Description

An **advanced and configurable file walker**.

It walks a tree in a filesystem and outputs the paths to a channel for asynchronous pipelines.

## Installation

```bash
go get github.com/mpragliola/walker
```

## Usage

### Builder

You **can't instantiate walkers directly**; use
the **builder**. The builder pattern will allow you to define your walker
with a fluent interface, initialize it, add filters,
and build it also in different steps.

### How to

In its simplest form you would create a walker and then make it walk a filesystem:

```go
// Walk everything under myRoot/ folder
w := walker.NewBuilder("myRoot").Build()
pathsChannel, err := w.Walk()
```

It will return a **channel** that you can consume asynchronously.

You can read it into a slice by using the helper function
`walker.Sink()`:

```go
var paths []string
paths = walker.Sink(pathsChannel)
```

### A more complex example

In this example we will
* walk folder `example/`
* extract only the Markdown files
* but excluding files that begin with underscore (note the use of `Base()` to filter by filename only):

```golang
ch, err := walker.NewBuilder("example").
    AddFilters(
        walker.FilterExtensions(".md"),
        walker.Not(walker.Base(walker.FilterRegex("^_.*")),
    ).
    Build().
    Walk()
```

### Use another filesystem

If you use other filesystems (f. ex. mocking FS for testing
purpose), if they respect the `fs.FS` interface you can specify them so:

```go
// Create a test mock filesystem
myFS := fstest.MapFS{
  "file1": {},
  "folder1/file2": {},
}

// Make the walker walk that filesystem
w := walker.NewBuilder("myRoot").
    SetFS(myFS).
    Build()
```

### Simple filtering

Add **filters** to restrict the search. A filter is of
type `walter.WalkFilter`:

```go 
textFilter := walker.FilterExtensions(".txt", ".md")
underscoreStartFilter := walker.FilterRegex("^_.*")

// Walk text and Markdown files that start with underscore
w := walker.NewBuilder("myRoot").
    AddFilters(
        textFilter,
        walker.Base(underscoreStartFilter)
    ),
    Build()
```

#### Built in filters

* `FilterIdentity()` does nothing, used only for testing
* `FilterExtensions(exts ...string)` will filter for the
  specified extensions:
* `FilterRegex(pattern string)` as the name implies will
  match for a regular expression

#### Custom filter

You can implement any filter as long as it matches
`walker.WalkFilter`'s signature:

```go
type WalkFilter func(path string) (bool, error)
```

Example:

```go
noCigsFlt := func(path string) (bool, error) {
   return path != "cigarette", nil 
}

noCigsBaseFlt := walker.Base(noCigsFlt)
```

#### Filter transformers

Filter transformers wrap one filter and return another
with a changed behavior.

* `walker.Base()` passes only the basename of the file
  to the wrapped filter
* `walker.Not()` will negate the underlying filter

They can be combined:

```go
// only paths whose basename does NOT begin with "_"
flt := walker.Not(walker.Base(walker.FilterRegex("^_.*")))
```