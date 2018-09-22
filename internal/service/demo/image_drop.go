package demo

import (
	"fmt"
	"os"

	"github.com/medtune/beta-platform/internal/jsonutil"
)

// DropImagePath .
func DropImagePath(p string) error {
	if p == "" {
		return fmt.Errorf("got empty arg: %v", p)
	}

	err := os.Remove("./" + p)
	if err != nil {
		return err
	}
	return nil
}

// DropImage .
func DropImage(infData *jsonutil.RunImageInference) error {
	return DropImagePath(infData.File)
}
