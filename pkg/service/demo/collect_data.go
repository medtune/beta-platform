package demo

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

// CollectImagesData collect pictures informations for demo
func CollectImagesData(demo string) ([]data.Image, error) {
	path := fmt.Sprintf("static/demos/%v/images", demo)
	if _, err := os.Stat(path); err != nil || os.IsNotExist(err) {
		return nil, err
	}

	// ls dir
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	images := make([]data.Image, 0, 0)

	// walk over file and collect images
	for _, f := range files {
		if p := strings.Split(f.Name(), "_"); len(p) > 2 {
			continue
		}
		// ignore 'debug.*' and sub directories
		if basename := strings.Split(f.Name(), ".")[0]; basename != "debug" && !f.IsDir() {
			images = append(images, data.Image{
				Name:     basename,
				Filename: f.Name(),
			})
		}
	}

	return images, nil
}
