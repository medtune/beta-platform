package demos

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

// CollectImagesData collect pictures informations for demo
func CollectImagesData(demo string) ([]data.Image, error) {
	path := fmt.Sprintf("static/demos/%v/images", demo)
	if _, err := os.Stat(path); err != nil || os.IsNotExist(err) {
		return nil, err
	}
	images := make([]data.Image, 0, 0)

	// ls dir
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	// walk over file
	for _, f := range files {
		if !f.IsDir() {
			images = append(images, data.Image{
				Name: strings.Split(f.Name(), ".")[0],
				URL:  "/" + path + "/" + f.Name(),
			})
		}
	}

	return images, nil
}

// DropImage .
func DropImage(infData *jsonutil.RunImageInference) error {
	if infData.File == "" {
		return fmt.Errorf("file field is empty: got struct %v", infData)
	}

	err := os.Remove("." + infData.File)
	if err != nil {
		return err
	}

	return nil
}

// DropImagePath .
func DropImagePath(p string) error {
	if p == "" {
		return fmt.Errorf("file field is empty: got %v", p)
	}

	err := os.Remove("./" + p)
	if err != nil {
		return err
	}
	return nil
}
