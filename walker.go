package walker

import (
	"io/fs"
)

type walker struct {
	fs      fs.FS
	root    string
	filters []WalkFilter
}

func (w *walker) Walk() (chan string, error) {
	ch := make(chan string)

	go func() {
		defer close(ch)

		err := fs.WalkDir(w.fs, w.root, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.IsDir() {
				return nil
			}

			for _, filter := range w.filters {
				ok, err := filter(path)

				if err != nil {
					return err
				} else if !ok {
					return nil
				}

			}

			ch <- path

			return nil

		})

		if err != nil {
			panic(err)
		}
	}()

	return ch, nil
}
