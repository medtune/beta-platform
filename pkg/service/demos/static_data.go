package demos

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
)

// Collect pictures informations for demo
func CollectImagesData(demo string) ([]data.Image, error) {
	path := "static/demos/" + demo + "/images"
	if _, err := os.Stat(path); err != nil || os.IsNotExist(err) {
		return nil, err
	}

	images := make([]data.Image, 0, 0)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if !f.IsDir() {
			images = append(images, data.Image{
				Name: f.Name(),
				URL:  "/" + path + "/" + f.Name(),
			})
		}
	}

	return images, nil
}

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
