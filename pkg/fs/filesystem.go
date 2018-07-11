package fs

import "io"

type commander struct {
	workdir string
}

func New() *commander {
	return &commander{workdir: "."}
}

func (*commander) Move(file string, dest string) error {
	return nil
}

func (*commander) Create(body io.ReadCloser, file string, dest string) error {
	return nil
}

func (*commander) Delete(file string) error {
	return nil
}
