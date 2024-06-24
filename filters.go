package walker

import (
	"path/filepath"
	"regexp"
	"strings"
)

type WalkFilter func(path string) (bool, error)

// -------------------
// Filter transformers
// -------------------

// Not returns the negation of the given filter.
func Not(filter WalkFilter) WalkFilter {
	return func(path string) (bool, error) {
		ok, err := filter(path)
		return !ok, err
	}
}

// Base returns a filter that applies the given filter to the base name of the path.
func Base(filter WalkFilter) WalkFilter {
	return func(path string) (bool, error) {
		return filter(filepath.Base(path))
	}
}

// -------
// Filters
// -------

// FilterIdentity returns a filter that matches all paths.
func FilterIdentity() WalkFilter {
	return func(path string) (bool, error) {
		return true, nil
	}
}

// FilterExtensions returns a filter that matches the given extensions.
func FilterExtensions(exts ...string) WalkFilter {
	return func(path string) (bool, error) {
		for _, ext := range exts {
			if filepath.Ext(path) == ext {
				return true, nil
			}
		}

		return false, nil
	}
}

// FilterRegex returns a filter that matches the given regular expression.
func FilterRegex(pattern string) WalkFilter {
	return func(path string) (bool, error) {
		re, err := regexp.Compile(pattern)
		if err != nil {
			return false, err
		}

		return re.MatchString(path), nil
	}
}

// FilterStartsWith returns a filter that matches paths that
// start with the given prefix.
func FilterStartsWith(prefix string) WalkFilter {
	return func(path string) (bool, error) {
		return strings.HasPrefix(path, prefix), nil
	}
}
